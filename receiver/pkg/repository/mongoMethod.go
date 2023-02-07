package repository

import (
	"context"
	"fmt"
	"github.com/sinyavcev/microservices/receiver/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	collection *mongo.Collection
}

func (d *DB) Create(ctx context.Context, stock model.Stock) (string, error) {
	result, err := d.collection.InsertOne(ctx, stock)
	if err != nil {
		return "", fmt.Errorf("failed to create stock due to error: %v", err)
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convert object id to hex")
	}
	return oid.Hex(), nil
}

func (d *DB) FindAllBySymbol(ctx context.Context, symbol string) (stocks []*model.Stock, err error) {
	filter := bson.M{"symbol": symbol}
	result, err := d.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = result.Decode(&stocks)
	if err != nil {
		return nil, err
	}

	return stocks, nil
}

func NewMongoMethod(collection *mongo.Collection) *DB {
	return &DB{collection: collection}
}
