package mongodb

import (
	"context"

	"github.com/utkarsh17ife/infinity/product"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	db *mongo.Client
}

// New ....
func New(db *mongo.Client) product.Repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) CreateProduct(ctx context.Context, product *product.Product) (*product.Product, error) {
	return nil, nil
}
func (repo *repository) GetProductByID(ctx context.Context, productID string) (*product.Product, error) {
	return nil, nil
}
func (repo *repository) GetProductByLocation(ctx context.Context, location string) (*product.Product, error) {
	return nil, nil
}
func (repo *repository) GetProductByCreatorID(ctx context.Context, creatorID string) ([]*product.Product, error) {
	return nil, nil
}
func (repo *repository) GetProductList(ctx context.Context) ([]*product.Product, error) {
	return nil, nil
}
func (repo *repository) UpdateProduct(ctx context.Context, updateProduct *product.Product) (*product.Product, error) {
	return nil, nil
}
func (repo *repository) SearchProduct(ctx context.Context, searchString string) ([]*product.Product, error) {
	return nil, nil
}
func (repo *repository) RemoveProduct(ctx context.Context, productID string) error {
	return nil
}
