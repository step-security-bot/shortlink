package leveldb

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"github.com/syndtr/goleveldb/leveldb"

	"github.com/batazor/shortlink/internal/store/query"
	"github.com/batazor/shortlink/pkg/link"
)

// LevelDBConfig ...
type LevelDBConfig struct { // nolint unused
	Path string
}

// LevelDBLinkList implementation of store interface
type LevelDBLinkList struct { // nolint unused
	ctx context.Context

	client *leveldb.DB
	config LevelDBConfig
}

// Init ...
func (l *LevelDBLinkList) Init(ctx context.Context) error {
	var err error

	// Set context
	l.ctx = ctx

	// Set configuration
	l.setConfig()

	l.client, err = leveldb.OpenFile("/tmp/links.db", nil)
	if err != nil {
		return err
	}

	return nil
}

// Close ...
func (l *LevelDBLinkList) Close() error {
	return l.client.Close()
}

// Migrate ...
func (l *LevelDBLinkList) migrate() error { // nolint unused
	return nil
}

// Add ...
func (l *LevelDBLinkList) Add(ctx context.Context, source *link.Link) (*link.Link, error) {
	data, err := link.NewURL(source.Url) // Create a new link
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = l.client.Put([]byte(data.Hash), payload, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Get ...
func (l *LevelDBLinkList) Get(ctx context.Context, id string) (*link.Link, error) {
	value, err := l.client.Get([]byte(id), nil)
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	var response link.Link

	err = json.Unmarshal(value, &response)
	if err != nil {
		return nil, err
	}

	if response.Url == "" {
		return nil, &link.NotFoundError{Link: &link.Link{Url: id}, Err: fmt.Errorf("Not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (l *LevelDBLinkList) List(ctx context.Context, filter *query.Filter) ([]*link.Link, error) { // nolint unused
	links := []*link.Link{}
	iterator := l.client.NewIterator(nil, nil)

	for iterator.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		value := iterator.Value()

		var response link.Link

		err := json.Unmarshal(value, &response)
		if err != nil {
			return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
		}

		links = append(links, &response)
	}

	iterator.Release()
	err := iterator.Error()
	if err != nil {
		return nil, &link.NotFoundError{Link: &link.Link{}, Err: fmt.Errorf("Not found links")}
	}

	return links, nil
}

// Update ...
func (l *LevelDBLinkList) Update(ctx context.Context, data *link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (l *LevelDBLinkList) Delete(ctx context.Context, id string) error {
	err := l.client.Delete([]byte(id), nil)
	return err
}

// setConfig - set configuration
func (l *LevelDBLinkList) setConfig() {
	viper.AutomaticEnv()
	viper.SetDefault("STORE_LEVELDB_PATH", "/tmp/links.db")

	l.config = LevelDBConfig{
		Path: viper.GetString("STORE_LEVELDB_PATH"),
	}
}
