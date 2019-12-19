package implementation

import (
	"context"
	"log"

	"github.com/utkarsh17ife/infinity/product"
)

// service implements the Product Service
type service struct {
	repository product.Repository
	logger     log.Logger
}

// NewService creates and returns a new Order service instance
func NewService(rep product.Repository) product.Service {
	return &service{
		repository: rep,
	}
}

func (s *service) CreateProduct(ctx context.Context, product *product.Product) (*product.Product, error) {
	return nil, nil
}
func (s *service) GetProductByID(ctx context.Context, productID string) (*product.Product, error) {
	return nil, nil
}
func (s *service) GetProductByLocation(ctx context.Context, location string) (*product.Product, error) {
	return nil, nil
}
func (s *service) GetProductByCreatorID(ctx context.Context, creatorID string) ([]*product.Product, error) {
	return nil, nil
}
func (s *service) GetProductList(ctx context.Context) ([]*product.Product, error) {
	return nil, nil
}
func (s *service) UpdateProduct(ctx context.Context, updateProduct *product.Product) (*product.Product, error) {
	return nil, nil
}
func (s *service) SearchProduct(ctx context.Context, searchString string) ([]*product.Product, error) {
	return nil, nil
}
func (s *service) RemoveProduct(ctx context.Context, productID string) error {
	return nil
}
