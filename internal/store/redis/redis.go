package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"

	"github.com/go-redis/redis"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

// RedisConfig ...
type RedisConfig struct { // nolint unused
	URI string
}

// RedisLinkList implementation of store interface
type RedisLinkList struct { // nolint unused
	ctx context.Context

	client *redis.Client
	config RedisConfig
}

// Init ...
func (r *RedisLinkList) Init(ctx context.Context) error {
	// Set context
	r.ctx = ctx

	// Set configuration
	r.setConfig()

	// Connect to Redis
	r.client = redis.NewClient(&redis.Options{
		Addr:     r.config.URI,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := r.client.Ping().Result(); err != nil {
		return err
	}

	return nil
}

// Close ...
func (r *RedisLinkList) Close() error {
	return r.client.Close()
}

// Migrate ...
func (r *RedisLinkList) migrate() error { // nolint unused
	return nil
}

// Get ...
func (r *RedisLinkList) Get(ctx context.Context, id string) (*link.Link, error) {
	val, err := r.client.Get(id).Result()
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link

	if err = json.Unmarshal([]byte(val), &response); err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed parse link: %s", id)}
	}

	return &response, nil
}

// List ...
func (r *RedisLinkList) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) { // nolint unused
	keys := r.client.Keys("*")
	links := []*link.Link{}

	for _, key := range keys.Val() {
		var response link.Link
		val, err := r.client.Get(key).Result()
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		if err = json.Unmarshal([]byte(val), &response); err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		links = append(links, &response)
	}

	return links, nil
}

// Add ...
func (r *RedisLinkList) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	val, err := json.Marshal(data)
	if err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed marsharing link: %s", data.Url)}
	}

	if err = r.client.Set(data.Hash, val, 0).Err(); err != nil {
		return nil, &link.NotFoundError{Link: data, Err: fmt.Errorf("Failed save link: %s", data.Url)}
	}

	return data, nil
}

// Update ...
func (r *RedisLinkList) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (r *RedisLinkList) Delete(ctx context.Context, id string) error {
	if err := r.client.Del(id).Err(); err != nil {
		return &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Failed save link: %s", id)}
	}

	return nil
}

// setConfig - set configuration
func (r *RedisLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_REDIS_URI", "localhost:6379")

	r.config = RedisConfig{
		URI: viper.GetString("STORE_REDIS_URI"),
	}
}
