package market

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// service implements the Item Service
type service struct {
	repository Repository
	logger     log.Logger
}

// NewService creates and returns a new Order service instance
func NewService(rep Repository) Service {
	return &service{
		repository: rep,
	}
}

func (s *service) CreateItem(ctx context.Context, item *Item) (*Item, error) {
	item, err := s.repository.CreateItem(ctx, item)
	return item, err
}
func (s *service) GetItemByID(ctx context.Context, itemID primitive.ObjectID) (*Item, error) {
	item, err := s.repository.GetItemByID(ctx, itemID)
	return item, err
}
func (s *service) GetItemByLocation(ctx context.Context, location string) (*Item, error) {
	return nil, nil
}
func (s *service) GetItemsByCreatorID(ctx context.Context, creatorID primitive.ObjectID) ([]*Item, error) {
	items, err := s.repository.GetItemsByCreatorID(ctx, creatorID)
	return items, err
}
func (s *service) GetItemList(ctx context.Context) ([]*Item, error) {
	return nil, nil
}
func (s *service) UpdateItem(ctx context.Context, updateItem *Item) (*Item, error) {
	return nil, nil
}
func (s *service) SearchItem(ctx context.Context, searchString string) ([]*Item, error) {
	return nil, nil
}
func (s *service) RemoveItem(ctx context.Context, itemID primitive.ObjectID) error {
	return nil
}
func (s *service) IncreaseItemQuantity(ctx context.Context, number int32) error {
	return nil
}
func (s *service) DecreaseItemQuantity(ctx context.Context, number int32) error {
	return nil
}
func (s *service) CheckQuantityAvaialble(ctx context.Context, amount int32) (bool, error) {
	return false, nil
}
