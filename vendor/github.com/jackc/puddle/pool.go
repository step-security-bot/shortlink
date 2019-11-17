package puddle

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	resourceStatusConstructing = 0
	resourceStatusIdle         = iota
	resourceStatusAcquired     = iota
	resourceStatusHijacked     = iota
)

// ErrClosedPool occurs on an attempt to acquire a connection from a closed pool
// or a pool that is closed while the acquire is waiting.
var ErrClosedPool = errors.New("closed pool")

// Constructor is a function called by the pool to construct a resource.
type Constructor func(ctx context.Context) (res interface{}, err error)

// Destructor is a function called by the pool to destroy a resource.
type Destructor func(res interface{})

// Resource is the resource handle returned by acquiring from the pool.
type Resource struct {
	value        interface{}
	pool         *Pool
	creationTime time.Time
	status       byte
}

// Value returns the resource value.
func (res *Resource) Value() interface{} {
	if !(res.status == resourceStatusAcquired || res.status == resourceStatusHijacked) {
		panic("tried to access resource that is not acquired or hijacked")
	}
	return res.value
}

// Release returns the resource to the pool. res must not be subsequently used.
func (res *Resource) Release() {
	if res.status != resourceStatusAcquired {
		panic("tried to release resource that is not acquired")
	}
	res.pool.releaseAcquiredResource(res)
}

// Destroy returns the resource to the pool for destruction. res must not be
// subsequently used.
func (res *Resource) Destroy() {
	if res.status != resourceStatusAcquired {
		panic("tried to destroy resource that is not acquired")
	}
	res.pool.destroyAcquiredResource(res)
}

// Hijack assumes ownership of the resource from the pool. Caller is responsible
// for cleanup of resource value.
func (res *Resource) Hijack() {
	if res.status != resourceStatusAcquired {
		panic("tried to hijack resource that is not acquired")
	}
	res.pool.hijackAcquiredResource(res)
}

// CreationTime returns when the resource was created by the pool.
func (res *Resource) CreationTime() time.Time {
	if !(res.status == resourceStatusAcquired || res.status == resourceStatusHijacked) {
		panic("tried to access resource that is not acquired or hijacked")
	}
	return res.creationTime
}

// Pool is a concurrency-safe resource pool.
type Pool struct {
	cond       *sync.Cond
	destructWG *sync.WaitGroup

	allResources  []*Resource
	idleResources []*Resource

	constructor Constructor
	destructor  Destructor
	maxSize     int32

	acquireCount         int64
	acquireDuration      time.Duration
	emptyAcquireCount    int64
	canceledAcquireCount int64

	closed bool
}

// NewPool creates a new pool. Panics if maxSize is less than 1.
func NewPool(constructor Constructor, destructor Destructor, maxSize int32) *Pool {
	if maxSize < 1 {
		panic("maxSize is less than 1")
	}

	return &Pool{
		cond:        sync.NewCond(new(sync.Mutex)),
		destructWG:  &sync.WaitGroup{},
		maxSize:     maxSize,
		constructor: constructor,
		destructor:  destructor,
	}
}

// Close destroys all resources in the pool and rejects future Acquire calls.
// Blocks until all resources are returned to pool and destroyed.
func (p *Pool) Close() {
	p.cond.L.Lock()
	p.closed = true

	for _, res := range p.idleResources {
		p.allResources = removeResource(p.allResources, res)
		go p.destructResourceValue(res.value)
	}
	p.idleResources = nil
	p.cond.L.Unlock()

	// Wake up all go routines waiting for a resource to be returned so they can terminate.
	p.cond.Broadcast()

	p.destructWG.Wait()
}

// Stat is a snapshot of Pool statistics.
type Stat struct {
	constructingResources int32
	acquiredResources     int32
	idleResources         int32
	maxResources          int32
	acquireCount          int64
	acquireDuration       time.Duration
	emptyAcquireCount     int64
	canceledAcquireCount  int64
}

// TotalResource returns the total number of resources.
func (s *Stat) TotalResources() int32 {
	return s.constructingResources + s.acquiredResources + s.idleResources
}

// ConstructingResources returns the number of resources with construction in progress in
// the pool.
func (s *Stat) ConstructingResources() int32 {
	return s.constructingResources
}

// AcquiredResources returns the number of acquired resources in the pool.
func (s *Stat) AcquiredResources() int32 {
	return s.acquiredResources
}

// IdleResources returns the number of idle resources in the pool.
func (s *Stat) IdleResources() int32 {
	return s.idleResources
}

// MaxResources returns the maximum size of the pool.
func (s *Stat) MaxResources() int32 {
	return s.maxResources
}

// AcquireCount returns the number of successful acquires from the pool.
func (s *Stat) AcquireCount() int64 {
	return s.acquireCount
}

// AcquireDuration returns the total duration of all successful acquires from
// the pool.
func (s *Stat) AcquireDuration() time.Duration {
	return s.acquireDuration
}

// EmptyAcquireCount returns the number of successful acquires from the pool
// that waited for a resource to be released or constructed because the pool was
// empty.
func (s *Stat) EmptyAcquireCount() int64 {
	return s.emptyAcquireCount
}

// CanceledAcquireCount returns the number of acquires from the pool
// that were canceled by a context.
func (s *Stat) CanceledAcquireCount() int64 {
	return s.canceledAcquireCount
}

// Stat returns the current pool statistics.
func (p *Pool) Stat() *Stat {
	p.cond.L.Lock()
	s := &Stat{
		maxResources:         p.maxSize,
		acquireCount:         p.acquireCount,
		emptyAcquireCount:    p.emptyAcquireCount,
		canceledAcquireCount: p.canceledAcquireCount,
		acquireDuration:      p.acquireDuration,
	}

	for _, res := range p.allResources {
		switch res.status {
		case resourceStatusConstructing:
			s.constructingResources += 1
		case resourceStatusIdle:
			s.idleResources += 1
		case resourceStatusAcquired:
			s.acquiredResources += 1
		}
	}

	p.cond.L.Unlock()
	return s
}

// Acquire gets a resource from the pool. If no resources are available and the pool
// is not at maximum capacity it will create a new resource. If the pool is at
// maximum capacity it will block until a resource is available. ctx can be used
// to cancel the Acquire.
func (p *Pool) Acquire(ctx context.Context) (*Resource, error) {
	startTime := time.Now()
	p.cond.L.Lock()
	if doneChan := ctx.Done(); doneChan != nil {
		select {
		case <-ctx.Done():
			p.canceledAcquireCount += 1
			p.cond.L.Unlock()
			return nil, ctx.Err()
		default:
		}
	}

	emptyAcquire := false

	for {
		if p.closed {
			p.cond.L.Unlock()
			return nil, ErrClosedPool
		}

		// If a resource is available now
		if len(p.idleResources) > 0 {
			res := p.idleResources[len(p.idleResources)-1]
			p.idleResources = p.idleResources[:len(p.idleResources)-1]
			res.status = resourceStatusAcquired
			if emptyAcquire {
				p.emptyAcquireCount += 1
			}
			p.acquireCount += 1
			p.acquireDuration += time.Now().Sub(startTime)
			p.cond.L.Unlock()
			return res, nil
		}

		emptyAcquire = true

		// If there is room to create a resource do so
		if len(p.allResources) < int(p.maxSize) {
			res := &Resource{pool: p, creationTime: startTime, status: resourceStatusConstructing}
			p.allResources = append(p.allResources, res)
			p.destructWG.Add(1)
			p.cond.L.Unlock()

			value, err := p.constructResourceValue(ctx)
			p.cond.L.Lock()
			if err != nil {
				p.allResources = removeResource(p.allResources, res)
				p.destructWG.Done()

				select {
				case <-ctx.Done():
					if err == ctx.Err() {
						p.canceledAcquireCount += 1
					}
				default:
				}

				p.cond.L.Unlock()
				return nil, err
			}

			res.value = value
			res.status = resourceStatusAcquired
			p.emptyAcquireCount += 1
			p.acquireCount += 1
			p.acquireDuration += time.Now().Sub(startTime)
			p.cond.L.Unlock()
			return res, nil
		}

		if ctx.Done() == nil {
			p.cond.Wait()
		} else {
			// Convert p.cond.Wait into a channel
			waitChan := make(chan struct{}, 1)
			go func() {
				p.cond.Wait()
				waitChan <- struct{}{}
			}()

			select {
			case <-ctx.Done():
				p.cond.L.Lock()
				p.canceledAcquireCount += 1
				p.cond.L.Unlock()

				// Allow goroutine waiting for signal to exit. Re-signal since we couldn't
				// do anything with it. Another goroutine might be waiting.
				go func() {
					<-waitChan
					p.cond.Signal()
					p.cond.L.Unlock()
				}()

				return nil, ctx.Err()
			case <-waitChan:
			}
		}
	}
}

// AcquireAllIdle atomically acquires all currently idle resources. Its intended
// use is for health check and keep-alive functionality. It does not update pool
// statistics.
func (p *Pool) AcquireAllIdle() []*Resource {
	p.cond.L.Lock()

	for _, res := range p.idleResources {
		res.status = resourceStatusAcquired
	}
	resources := make([]*Resource, len(p.idleResources))
	copy(resources, p.idleResources)
	p.idleResources = p.idleResources[0:0]

	p.cond.L.Unlock()
	return resources
}

// releaseAcquiredResource returns res to the the pool.
func (p *Pool) releaseAcquiredResource(res *Resource) {
	p.cond.L.Lock()

	if !p.closed {
		res.status = resourceStatusIdle
		p.idleResources = append(p.idleResources, res)
	} else {
		p.allResources = removeResource(p.allResources, res)
		go p.destructResourceValue(res.value)
	}

	p.cond.L.Unlock()
	p.cond.Signal()
}

// Remove removes res from the pool and closes it. If res is not part of the
// pool Remove will panic.
func (p *Pool) destroyAcquiredResource(res *Resource) {
	p.cond.L.Lock()

	p.allResources = removeResource(p.allResources, res)
	go p.destructResourceValue(res.value)

	p.cond.L.Unlock()
	p.cond.Signal()
}

func (p *Pool) hijackAcquiredResource(res *Resource) {
	p.cond.L.Lock()

	p.allResources = removeResource(p.allResources, res)
	res.status = resourceStatusHijacked
	p.destructWG.Done() // not responsible for destructing hijacked resources

	p.cond.L.Unlock()
	p.cond.Signal()
}

func removeResource(slice []*Resource, res *Resource) []*Resource {
	for i := range slice {
		if slice[i] == res {
			slice[i] = slice[len(slice)-1]
			return slice[:len(slice)-1]
		}
	}

	panic("BUG: removeResource could not find res in slice")
}

func (p *Pool) constructResourceValue(ctx context.Context) (interface{}, error) {
	return p.constructor(ctx)
}

func (p *Pool) destructResourceValue(value interface{}) {
	p.destructor(value)
	p.destructWG.Done()
}
