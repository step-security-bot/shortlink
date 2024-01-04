package order

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/segmentio/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"

	order_application "github.com/shortlink-org/shortlink/internal/boundaries/payment/billing/application/order"
	billing "github.com/shortlink-org/shortlink/internal/boundaries/payment/billing/domain/billing/order/v1"
)

type API struct {
	jsonpb protojson.MarshalOptions //nolint:structcheck // ignore

	orderService *order_application.OrderService
}

func New(orderService *order_application.OrderService) (*API, error) {
	return &API{
		orderService: orderService,
	}, nil
}

// Routes creates a REST router
func (api *API) Routes(r chi.Router) {
	r.Get("/order/{hash}", api.get)
	r.Get("/orders", api.list)
	r.Post("/order", api.add)
	r.Delete("/order/{hash}", api.delete)
}

// Add - add
func (api *API) add(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	// Parse request
	var request billing.Order
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck,goconst // ignore

		return
	}

	newOrder, err := api.orderService.Add(r.Context(), &request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	res, err := api.jsonpb.Marshal(newOrder)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`)) //nolint:errcheck // ignore

		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(res) //nolint:errcheck // ignore
}

// Get - get
func (api *API) get(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
}

// List - list
func (api *API) list(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusOK)
}

// Delete - delete
func (api *API) delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	w.WriteHeader(http.StatusNoContent)
}
