package market

import (
	"context"
	"errors"
)

// ErrInvalidArgument ....
var ErrInvalidArgument = errors.New("invalid argument")

// Service ....
type Service interface {
	CreateItem(ctx context.Context, item *Item) (*Item, error)
	GetItemByID(ctx context.Context, itemID string) (*Item, error)
	GetItemByLocation(ctx context.Context, location string) (*Item, error)
	GetItemByCreatorID(ctx context.Context, creatorID string) ([]*Item, error)
	GetItemList(ctx context.Context) ([]*Item, error)
	UpdateItem(ctx context.Context, updateItem *Item) (*Item, error)
	SearchItem(ctx context.Context, searchString string) ([]*Item, error)
	RemoveItem(ctx context.Context, itemID string) error

	IncreaseItemQuantity(ctx context.Context, number int32) error
	DecreaseItemQuantity(ctx context.Context, number int32) error
	CheckQuantityAvaialble(ctx context.Context, amount int32) (bool, error)
}
