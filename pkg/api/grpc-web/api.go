package grpcweb

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/batazor/shortlink/internal/notify"
	api_type "github.com/batazor/shortlink/pkg/api/type"
	"github.com/batazor/shortlink/pkg/link"
)

// GetLink ...
func (api *API) GetLink(ctx context.Context, req *link.Link) (*link.Link, error) {
	responseCh := make(chan interface{})

	go notify.Publish(api_type.METHOD_GET, req.Hash, responseCh, "RESPONSE_STORE_GET")

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return nil, err
		}
		response := r.Payload.(*link.Link) // nolint errcheck
		return response, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	}
}

// GetLinks ...
func (api *API) GetLinks(ctx context.Context, req *link.Link) (*link.Links, error) {
	responseCh := make(chan interface{})

	go notify.Publish(api_type.METHOD_LIST, nil, responseCh, "RESPONSE_STORE_LIST")

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_LIST")
		return nil, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return nil, err
		}
		links := r.Payload.([]*link.Link) // nolint errcheck

		response := link.Links{}
		for key := range links {
			response.Link = append(response.Link, links[key])
		}

		return &response, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_GET")
		return nil, err
	}
}

// CreateLink ...
func (api *API) CreateLink(ctx context.Context, req *link.Link) (*link.Link, error) {
	responseCh := make(chan interface{})

	go notify.Publish(api_type.METHOD_ADD, *req, responseCh, "RESPONSE_STORE_ADD")

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		return nil, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return nil, err
		}
		response := r.Payload.(*link.Link) // nolint errcheck
		return response, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_ADD")
		return nil, err
	}
}

// DeleteLink ...
func (api *API) DeleteLink(ctx context.Context, req *link.Link) (*empty.Empty, error) {
	responseCh := make(chan interface{})

	go notify.Publish(api_type.METHOD_DELETE, req.Hash, responseCh, "RESPONSE_STORE_DELETE")

	c := <-responseCh
	switch r := c.(type) {
	case nil:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
		return &empty.Empty{}, err
	case notify.Response:
		err := r.Error
		if err != nil {
			return nil, err
		}
		return &empty.Empty{}, nil
	default:
		err := fmt.Errorf("Not found subscribe to event %s", "METHOD_DELETE")
		return &empty.Empty{}, err
	}
}
