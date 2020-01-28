package httpserver

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/utkarsh17ife/infinity/market/transport"

	"github.com/gorilla/mux"
)

// NewHTTPServer is the http transport layer
func NewHTTPServer(ctx context.Context, endpoints transport.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

	r.Methods("POST").Path("/item").Handler(httptransport.NewServer(
		endpoints.CreateItem,
		decodeCreateItemRequest,
		encodeCreateItemResponse,
	))

	r.Methods("GET").Path("/item/{item_id}").Handler(httptransport.NewServer(
		endpoints.GetItemByID,
		decodeGetItemByIDRequest,
		encodeGetItemByIDResponse,
	))

	r.Methods("GET").Path("/item/createdby/{created_by_id}").Handler(httptransport.NewServer(
		endpoints.GetItemsByCreatorID,
		decodeGetItemsByCreatorIDRequest,
		encodeGetItemsByCreatorIDResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
