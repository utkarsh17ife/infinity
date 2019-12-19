package product

import (
	"context"
	"errors"
)

// ErrInvalidArgument ....
var ErrInvalidArgument = errors.New("invalid argument")

// Service ....
type Service interface {
	CreateProduct(ctx context.Context, product *Product) (*Product, error)

	GetProductByID(ctx context.Context, productID string) (*Product, error)

	GetProductByLocation(ctx context.Context, location string) (*Product, error)

	GetProductByCreatorID(ctx context.Context, creatorID string) ([]*Product, error)

	GetProductList(ctx context.Context) ([]*Product, error)

	UpdateProduct(ctx context.Context, updateProduct *Product) (*Product, error)

	SearchProduct(ctx context.Context, searchString string) ([]*Product, error)

	RemoveProduct(ctx context.Context, productID string) error
}
