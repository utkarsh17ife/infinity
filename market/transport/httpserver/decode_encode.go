package httpserver

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/utkarsh17ife/infinity/market/transport"
)

func decodeCreateItemRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req *transport.CreateItemRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func encodeCreateItemResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeGetItemByIDRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	req := &transport.GetItemByIDRequest{
		ItemID: vars["item_id"],
	}
	return req, nil
}
func encodeGetItemByIDResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
