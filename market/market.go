package market

import (
	"context"
	"time"

	models "github.com/utkarsh17ife/infinity/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Quantity .....
type Quantity struct {
	Amount uint
	Unit   string
}

// Item .....
type Item struct {
	ID           primitive.ObjectID `json:"item_id"`
	Name         string             `json:"item_name"`
	Typpe        string             `json:"type"`
	Category     string             `json:"category"`
	Description  string             `json:"description"`
	Seasons      []models.Season    `json:"seasons"`
	IsOrganic    bool               `json:"is_organic"`
	FarmID       string             `json:"farm_id"`
	PricePerUnit float32            `json:"price"`
	Location     models.Location    `json:"location"`
	Quantity     Quantity           `json:"quantity"`
	CreatedByID  string             `json:"created_by_id"`
	Message      string             `json:"message"`
	UpdatedAt    time.Time          `json:"updated_at"`
	CreatedAt    time.Time          `json:"created_at"`
}

// Repository .....
type Repository interface {
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
