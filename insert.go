package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOne(collection *mongo.Collection, data interface{}) (*mongo.InsertOneResult, error) {

	result, err := collection.InsertOne(context.TODO(), data)
	return result, err
}

func InsertMany(collection *mongo.Collection, data []interface{}) (*mongo.InsertManyResult, error) {

	result, err := collection.InsertMany(context.TODO(), data)
	return result, err
}
