package repository

import (
	"boilerplate-go/internal/domain/model"
	"context"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
    db *gorm.DB
}

func NewMySQLProductRepository(db *gorm.DB) ProductRepository {
    return &mysqlProductRepository{db: db}
}

func (r *mysqlProductRepository) CreateProduct(ctx context.Context, product *model.Product) error {
    return r.db.Create(product).Error
}

func (r *mysqlProductRepository) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
    var product model.Product
    if err := r.db.Where("id = ?", id).First(&product).Error; err != nil {
        return nil, err
    }
    return &product, nil
}
