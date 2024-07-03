package repository

import (
	"boilerplate-go/internal/domain/model"
	"context"
)

type ProductRepository interface {
    CreateProduct(ctx context.Context, product *model.Product) error
    GetProductByID(ctx context.Context, id string) (*model.Product, error)
}
