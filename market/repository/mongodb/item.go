package mongodb

import (
	"context"

	"github.com/utkarsh17ife/market/item"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Client
}

// New ....
func New(db *mongo.Client) item.Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) CreateItem(ctx context.Context, item *item.Item) (*item.Item, error) {
	return nil, nil
}
func (repo *repository) GetItemByID(ctx context.Context, itemID string) (*item.Item, error) {
	return nil, nil
}
func (repo *repository) GetItemByLocation(ctx context.Context, location string) (*item.Item, error) {
	return nil, nil
}
func (repo *repository) GetItemByCreatorID(ctx context.Context, creatorID string) ([]*item.Item, error) {
	return nil, nil
}
func (repo *repository) GetItemList(ctx context.Context) ([]*item.Item, error) {
	return nil, nil
}
func (repo *repository) UpdateItem(ctx context.Context, updateItem *item.Item) (*item.Item, error) {
	return nil, nil
}
func (repo *repository) SearchItem(ctx context.Context, searchString string) ([]*item.Item, error) {
	return nil, nil
}
func (repo *repository) RemoveItem(ctx context.Context, itemID string) error {
	return nil
}
func (repo *repository) IncreaseItemQuantity(ctx context.Context, number int32) error {
	return nil
}
func (repo *repository) DecreaseItemQuantity(ctx context.Context, number int32) error {
	return nil
}
func (repo *repository) CheckQuantityAvaialble(ctx context.Context, amount int32) (bool, error) {
	return nil, false
}
