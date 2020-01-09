package httpserver

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/utkarsh17ife/infinity/market/transport"
)

func decodeCreateItemRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req transport.CreateItemRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func encodeCreateItemResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
