package batch

import (
	"sync"
	"time"
)

// Batch is a structure for batch processing
type Batch struct {
	callback func([]*Item) interface{}
	items    []*Item
	interval time.Duration
	mu       sync.Mutex
}

// Item represents an item that can be pushed to the batch.
type Item struct {
	CallbackChannel chan interface{}
	Item            interface{}
}
