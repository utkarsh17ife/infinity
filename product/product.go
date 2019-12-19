package product

import (
	"context"
	"time"

	"github.com/utkarsh17ife/infinity/common_models"
)

// Quantity .....
type Quantity struct {
	Amount uint
	Unit   string
}

// Product .....
type Product struct {
	ID           string                 `json:"product_id"`
	Name         string                 `json:"product_name"`
	Typpe        string                 `json:"type"`
	Category     string                 `json:"category"`
	Description  string                 `json:"description"`
	Seasons      []common_models.Season `json:"seasons"`
	IsOrganic    bool                   `json:"is_organic"`
	FarmID       string                 `json:"farm_id"`
	PricePerUnit float32                `json:"price"`
	Location     common_models.Location `json:"location"`
	Quantity     Quantity               `json:"quantity"`
	CreatedByID  string                 `json:"created_by_id"`
	UpdatedAt    time.Time              `json:"updated_at"`
	CreatedAt    time.Time              `json:"created_at"`
}

// Repository .....
type Repository interface {
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	GetProductByID(ctx context.Context, productID string) (*Product, error)
	GetProductByLocation(ctx context.Context, location string) (*Product, error)
	GetProductByCreatorID(ctx context.Context, creatorID string) ([]*Product, error)
	GetProductList(ctx context.Context) ([]*Product, error)
	UpdateProduct(ctx context.Context, updateProduct *Product) (*Product, error)
	SearchProduct(ctx context.Context, searchString string) ([]*Product, error)
	RemoveProduct(ctx context.Context, productID string) error
}
