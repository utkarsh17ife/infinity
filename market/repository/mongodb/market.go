package mongodb

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/utkarsh17ife/infinity/market"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	NOT_FOUND = errors.New("Not found")
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

	item.ID = primitive.NewObjectID()
	res, err := collection.InsertOne(ctx, item)
	if err != nil {
		return nil, err
	}
	item.ID = res.InsertedID.(primitive.ObjectID)
	return item, nil
}

func (repo *repository) GetItemByID(ctx context.Context, _itemID primitive.ObjectID) (*market.Item, error) {

	var item *market.Item

	collection := repo.db.Database("cps_market").Collection("item")

	err := collection.FindOne(ctx, bson.M{"_id": _itemID}).Decode(&item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (repo *repository) GetItemByLocation(ctx context.Context, location string) (*market.Item, error) {
	return nil, nil
}
func (repo *repository) GetItemsByCreatorID(ctx context.Context, creatorID primitive.ObjectID) ([]*market.Item, error) {
	var items []*market.Item

	collection := repo.db.Database("cps_market").Collection("item")

	cur, err := collection.Find(ctx, bson.M{"created_by_id": creatorID})
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var item market.Item
		err := cur.Decode(&item)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, &item)
	}

	defer cur.Close(ctx)

	return items, nil
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
func (repo *repository) RemoveItem(ctx context.Context, itemID primitive.ObjectID) error {
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
