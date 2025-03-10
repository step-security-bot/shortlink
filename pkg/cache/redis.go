package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

// RedisClient wraps the Redis client with additional error handling
type RedisClient struct {
	client *redis.Client
}

// NewRedisClient creates a new Redis client wrapper
func NewRedisClient(opts *redis.Options) *RedisClient {
	return &RedisClient{
		client: redis.NewClient(opts),
	}
}

// Set stores a key-value pair with optional expiration
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return errors.NewCacheError("set", err)
	}
	return nil
}

// Get retrieves a value by key
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.ErrNotFound
	}
	if err != nil {
		return "", errors.NewCacheError("get", err)
	}
	return val, nil
}

// Delete removes a key
func (r *RedisClient) Delete(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return errors.NewCacheError("delete", err)
	}
	return nil
}

// Exists checks if a key exists
func (r *RedisClient) Exists(ctx context.Context, key string) (bool, error) {
	n, err := r.client.Exists(ctx, key).Result()
	if err != nil {
		return false, errors.NewCacheError("exists", err)
	}
	return n > 0, nil
}

// Close closes the Redis connection
func (r *RedisClient) Close() error {
	if err := r.client.Close(); err != nil {
		return errors.NewCacheError("close", err)
	}
	return nil
}

// Pipeline creates a new pipeline for batch operations
func (r *RedisClient) Pipeline() redis.Pipeliner {
	return r.client.Pipeline()
}

// Client returns the underlying Redis client
func (r *RedisClient) Client() *redis.Client {
	return r.client
}
