package service

import (
	"boilerplate-go/internal/domain/model"
	"boilerplate-go/internal/domain/repository"
	"boilerplate-go/internal/infrastructure/messaging"
	"context"
	"encoding/json"
)

type ProductService interface {
    CreateProduct(ctx context.Context, product *model.Product) error
    GetProductByID(ctx context.Context, id string) (*model.Product, error)
}

type productService struct {
    repo     repository.ProductRepository
    producer messaging.Producer
}

func NewProductService(repo repository.ProductRepository, producer messaging.Producer) ProductService {
    return &productService{
        repo:     repo,
        producer: producer,
    }
}

func (s *productService) CreateProduct(ctx context.Context, product *model.Product) error {
    body, err := json.Marshal(product)
    if err != nil {
        return err
    }

    err = s.producer.Publish("", "product.create", body)
    if err != nil {
        return err
    }
    return nil
}

func (s *productService) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
    return s.repo.GetProductByID(ctx, id)
}
