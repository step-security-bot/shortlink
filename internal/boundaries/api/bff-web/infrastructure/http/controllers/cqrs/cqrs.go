package cqrs

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"google.golang.org/protobuf/encoding/protojson"

	v1 "github.com/shortlink-org/shortlink/internal/boundaries/link/link/domain/link/v1"
	link_cqrs "github.com/shortlink-org/shortlink/internal/boundaries/link/link/infrastructure/rpc/cqrs/link/v1"
)

var jsonpb protojson.MarshalOptions

type LinkCQRSController struct {
	LinkCommandServiceClient link_cqrs.LinkCommandServiceClient
	LinkQueryServiceClient   link_cqrs.LinkQueryServiceClient
}

// GetLinkByCQRS - get link by hash
func (c *LinkCQRSController) GetLinkByCQRS(w http.ResponseWriter, r *http.Request, params any) {
	hash := chi.URLParam(r, "hash")
	if hash == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "need set hash URL"}`)) // nolint:errcheck

		return
	}

	response, err := c.LinkQueryServiceClient.Get(r.Context(), &link_cqrs.GetRequest{Hash: hash})
	var errorLink *v1.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := jsonpb.Marshal(response.Link)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint:errcheck
}

// GetLinksByCQRS - get links by filter
func (c *LinkCQRSController) GetLinksByCQRS(w http.ResponseWriter, r *http.Request, params any) {
	search := r.URL.Query().Get("search")
	response, err := c.LinkQueryServiceClient.List(r.Context(), &link_cqrs.ListRequest{Filter: search})
	var errorLink *v1.NotFoundError
	if errors.As(err, &errorLink) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	res, err := jsonpb.Marshal(response.Links)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) // nolint:errcheck

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res) // nolint:errcheck
}
