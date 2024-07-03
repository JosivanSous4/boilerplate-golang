package repository

import (
	"boilerplate-go/internal/domain/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoProductRepository struct {
    collection *mongo.Collection
}

func NewMongoProductRepository(db *mongo.Client, dbName, collectionName string) ProductRepository {
    collection := db.Database(dbName).Collection(collectionName)
    return &mongoProductRepository{collection: collection}
}

func (r *mongoProductRepository) CreateProduct(ctx context.Context, product *model.Product) error {
    _, err := r.collection.InsertOne(ctx, product)
    return err
}

func (r *mongoProductRepository) GetProductByID(ctx context.Context, id string) (*model.Product, error) {
    var product model.Product
    err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
    if err == mongo.ErrNoDocuments {
        return nil, errors.New("product not found")
    }
    return &product, err
}
