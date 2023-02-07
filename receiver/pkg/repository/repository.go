package repository

import (
	"context"
	"github.com/sinyavcev/microservices/receiver/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Stock interface {
	Create(ctx context.Context, stock model.Stock) (string, error)
	FindAllBySymbol(ctx context.Context, symbol string) (stocks []*model.Stock, err error)
}

type Repository struct {
	Stock
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		Stock: NewMongoMethod(db),
	}
}
