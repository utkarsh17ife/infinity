package mongodb

import (
	"context"

	"github.com/utkarsh17ife/infinity/market"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Client
}

// New ....
func New(db *mongo.Client) market.Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) CreateItem(ctx context.Context, item *market.Item) (*market.Item, error) {
	collection := repo.db.Database("cps_market").Collection("item")

	res, err := collection.InsertOne(context.Background(), item)
	if err != nil {
		return nil, err
	}
	item.ID = res.InsertedID.(primitive.ObjectID)

	return item, nil
}
func (repo *repository) GetItemByID(ctx context.Context, itemID string) (*market.Item, error) {
	return nil, nil
}
func (repo *repository) GetItemByLocation(ctx context.Context, location string) (*market.Item, error) {
	return nil, nil
}
func (repo *repository) GetItemByCreatorID(ctx context.Context, creatorID string) ([]*market.Item, error) {
	return nil, nil
}
func (repo *repository) GetItemList(ctx context.Context) ([]*market.Item, error) {
	return nil, nil
}
func (repo *repository) UpdateItem(ctx context.Context, updateItem *market.Item) (*market.Item, error) {
	return nil, nil
}
func (repo *repository) SearchItem(ctx context.Context, searchString string) ([]*market.Item, error) {
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
	return false, nil
}
