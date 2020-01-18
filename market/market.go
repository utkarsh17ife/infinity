package market

import (
	"context"
	"time"

	models "github.com/utkarsh17ife/infinity/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Quantity .....
type Quantity struct {
	Amount uint   `json:"amount"`
	Unit   string `json:"unit"`
}

// Item .....
type Item struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"item_name" bson:"item_name"`
	Typpe        string             `json:"type" bson:"type"`
	Category     string             `json:"category" bson:"category"`
	Description  string             `json:"description" bson:"description"`
	Seasons      []string           `json:"seasons" bson:"seasons"`
	IsOrganic    bool               `json:"is_organic" bson:"is_organic"`
	FarmID       primitive.ObjectID `json:"farm_id" bson:"farm_id"`
	PricePerUnit float32            `json:"price_per_unit" bson:"price_per_unit"`
	Location     models.Location    `json:"location" bson:"location"`
	Quantity     Quantity           `json:"quantity" bson:"quantity"`
	CreatedByID  string             `json:"created_by_id" bson:"created_by_id"`
	Message      string             `json:"message" bson:"message"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
}

// Repository .....
type Repository interface {
	CreateItem(ctx context.Context, item *Item) (*Item, error)
	GetItemByID(ctx context.Context, itemID primitive.ObjectID) (*Item, error)
	GetItemByLocation(ctx context.Context, location string) (*Item, error)
	GetItemByCreatorID(ctx context.Context, creatorID primitive.ObjectID) ([]*Item, error)
	GetItemList(ctx context.Context) ([]*Item, error)
	UpdateItem(ctx context.Context, updateItem *Item) (*Item, error)
	SearchItem(ctx context.Context, searchString string) ([]*Item, error)
	RemoveItem(ctx context.Context, itemID primitive.ObjectID) error
	IncreaseItemQuantity(ctx context.Context, number int32) error
	DecreaseItemQuantity(ctx context.Context, number int32) error
	CheckQuantityAvaialble(ctx context.Context, amount int32) (bool, error)
}
