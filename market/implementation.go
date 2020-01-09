package market

import (
	"context"
	"log"
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
func (s *service) GetItemByID(ctx context.Context, itemID string) (*Item, error) {
	return nil, nil
}
func (s *service) GetItemByLocation(ctx context.Context, location string) (*Item, error) {
	return nil, nil
}
func (s *service) GetItemByCreatorID(ctx context.Context, creatorID string) ([]*Item, error) {
	return nil, nil
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
func (s *service) RemoveItem(ctx context.Context, itemID string) error {
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
