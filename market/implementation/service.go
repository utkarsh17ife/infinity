package implementation

import (
	"context"
	"log"

	"github.com/utkarsh17ife/infinity/market"
)

// service implements the Item Service
type service struct {
	repository market.Repository
	logger     log.Logger
}

// NewService creates and returns a new Order service instance
func NewService(rep market.Repository) market.Service {
	return &service{
		repository: rep,
	}
}

func (s *service) CreateItem(ctx context.Context, item *market.Item) (*market.Item, error) {
	return nil, nil
}
func (s *service) GetItemByID(ctx context.Context, itemID string) (*market.Item, error) {
	return nil, nil
}
func (s *service) GetItemByLocation(ctx context.Context, location string) (*market.Item, error) {
	return nil, nil
}
func (s *service) GetItemByCreatorID(ctx context.Context, creatorID string) ([]*market.Item, error) {
	return nil, nil
}
func (s *service) GetItemList(ctx context.Context) ([]*market.Item, error) {
	return nil, nil
}
func (s *service) UpdateItem(ctx context.Context, updateItem *market.Item) (*market.Item, error) {
	return nil, nil
}
func (s *service) SearchItem(ctx context.Context, searchString string) ([]*market.Item, error) {
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
