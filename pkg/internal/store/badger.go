package store

import (
	"encoding/json"
	"fmt"

	"github.com/batazor/shortlink/pkg/link"
	"github.com/dgraph-io/badger"
)

// BadgerLinkList implementation of store interface
type BadgerLinkList struct { // nolint unused
	client *badger.DB
}

// Init ...
func (b *BadgerLinkList) Init() error {
	var err error
	b.client, err = badger.Open(badger.DefaultOptions("/tmp/links.badger"))
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func (b *BadgerLinkList) Get(id string) (*link.Link, error) {
	var valCopy []byte

	err := b.client.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(id))
		if err != nil {
			return err
		}

		err = item.Value(func(val []byte) error {
			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)

			return nil
		})
		if err != nil {
			return err
		}

		// Alternatively, you could also use item.ValueCopy().
		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	var response link.Link

	err = json.Unmarshal(valCopy, &response)
	if err != nil {
		return nil, err
	}

	if response.URL == "" {
		return nil, &link.NotFoundError{Link: link.Link{URL: id}, Err: fmt.Errorf("not found id: %s", id)}
	}

	return &response, nil
}

// List ...
func (b *BadgerLinkList) List() ([]*link.Link, error) {
	var list [][]byte

	err := b.client.View(func(txn *badger.Txn) error {
		iterator := txn.NewIterator(badger.DefaultIteratorOptions)
		defer iterator.Close()

		for iterator.Rewind(); iterator.Valid(); iterator.Next() {
			var valCopy []byte
			item := iterator.Item()

			err := item.Value(func(val []byte) error {
				// Copying or parsing val is valid.
				valCopy = append([]byte{}, val...)

				return nil
			})
			if err != nil {
				return err
			}

			// Alternatively, you could also use item.ValueCopy().
			valCopy, err = item.ValueCopy(nil)
			if err != nil {
				return err
			}

			list = append(list, valCopy)
		}

		return nil
	})
	if err != nil {
		return nil, &link.NotFoundError{Link: link.Link{}, Err: fmt.Errorf("not found links: %s", err)}
	}

	response := make([]*link.Link, len(list))

	for index, link := range list {
		err = json.Unmarshal(link, &response[index])
		if err != nil {
			return nil, err
		}
	}

	return response, nil
}

// Add ...
func (b *BadgerLinkList) Add(data link.Link) (*link.Link, error) {
	hash := data.CreateHash([]byte(data.URL), []byte("secret"))
	data.Hash = hash[:7]

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	err = b.client.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(data.Hash), payload)
	})
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Update ...
func (b *BadgerLinkList) Update(data link.Link) (*link.Link, error) {
	return nil, nil
}

// Delete ...
func (b *BadgerLinkList) Delete(id string) error {
	err := b.client.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(id))
		return err
	})
	return err
}
