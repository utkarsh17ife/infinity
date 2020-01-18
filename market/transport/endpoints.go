package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/utkarsh17ife/infinity/market"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Endpoints struct {
	CreateItem             endpoint.Endpoint
	GetItemByID            endpoint.Endpoint
	GetItemByLocation      endpoint.Endpoint
	GetItemByCreatorID     endpoint.Endpoint
	GetItemList            endpoint.Endpoint
	UpdateItem             endpoint.Endpoint
	SearchItem             endpoint.Endpoint
	RemoveItem             endpoint.Endpoint
	IncreaseItemQuantity   endpoint.Endpoint
	DecreaseItemQuantity   endpoint.Endpoint
	CheckQuantityAvaialble endpoint.Endpoint
}

func MakeEndpoints(s market.Service) Endpoints {
	return Endpoints{
		CreateItem:  makeCreateItemEndpoint(s),
		GetItemByID: makeGetItemByIDEndpoint(s),
	}
}

type CreateItemRequest struct {
	Item *market.Item `json:"item"`
}

type CreateItemResponse struct {
	Item    *market.Item `json:"item"`
	Message string       `json:"message"`
	Error   string       `json:"error"`
}

func makeCreateItemEndpoint(s market.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		var item *market.Item
		req := request.(*CreateItemRequest)
		item, err = s.CreateItem(ctx, req.Item)
		message := "Item is created"
		if err != nil {
			errorMessage := "Failed to create the item"
			return CreateItemResponse{Item: nil, Message: errorMessage, Error: err.Error()}, nil
		}
		return CreateItemResponse{Item: item, Message: message}, nil
	}
}

type GetItemByIDRequest struct {
	ItemID string
}

type GetItemByIDResponse struct {
	Item    *market.Item `json:"item"`
	Message string       `json:"message"`
	Error   string       `json:"error"`
}

func makeGetItemByIDEndpoint(s market.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var err error
		req := request.(*GetItemByIDRequest)

		itemID, err := primitive.ObjectIDFromHex(req.ItemID)

		if err != nil {
			errorMessage := "Invalid item id"
			return CreateItemResponse{Item: nil, Message: errorMessage, Error: err.Error()}, nil
		}

		item, err := s.GetItemByID(ctx, itemID)
		if err != nil {
			errorMessage := "Failed to get item"
			return CreateItemResponse{Item: nil, Message: errorMessage, Error: err.Error()}, nil
		}

		message := "Item found"
		return CreateItemResponse{Item: item, Message: message}, nil
	}
}
