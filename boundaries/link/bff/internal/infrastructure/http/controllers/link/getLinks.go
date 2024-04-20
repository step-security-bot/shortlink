package link

import (
	"net/http"

	"github.com/shortlink-org/shortlink/boundaries/link/bff/internal/infrastructure/http/api"
)

// GetLinks - get links
func (c *Controller) GetLinks(w http.ResponseWriter, r *http.Request, params api.GetLinksParams) {
	// TODO: add mapper for filter

	// request := &link_rpc.ListRequest{}
	//
	// if params.Cursor != nil {
	// 	request.Cursor = *params.Cursor
	// }
	//
	// if params.Limit != nil {
	// 	request.Limit = uint32(*params.Limit)
	// }
	//
	// result, err := c.linkServiceClient.List(r.Context(), request)
	// if err != nil {
	// 	var errorLink *v1.NotFoundError
	//
	// 	if errors.Is(err, errorLink) {
	// 		w.WriteHeader(http.StatusNotFound)
	// 		_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck
	//
	// 		return
	// 	}
	//
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	_ = json.NewEncoder(w).Encode(ErrMessages(err)) //nolint:errcheck
	//
	// 	return
	// }
	//
	// response := &api.GetLinks200Response{
	// 	Links:      make([]api.Link, 0, len(result.GetLinks().GetLink())),
	// 	NextCursor: result.GetCursor(),
	// }
	//
	// for _, link := range result.GetLinks().GetLink() {
	// 	response.Links = append(response.Links, api.Link{
	// 		Url:       link.GetUrl(),
	// 		Hash:      link.GetHash(),
	// 		Describe:  link.GetDescribe(),
	// 		CreatedAt: link.GetCreatedAt().AsTime(),
	// 		UpdatedAt: link.GetUpdatedAt().AsTime(),
	// 	})
	// }
	//
	// w.WriteHeader(http.StatusOK)
	// err = json.NewEncoder(w).Encode(response)
	// if err != nil {
	// 	c.log.Error(err.Error())
	// }
}
