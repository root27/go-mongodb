package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func FindOne(collection *mongo.Collection, filter interface{}) *mongo.SingleResult {

	result := collection.FindOne(context.TODO(), filter)

	return result
}

func FindMany(collection *mongo.Collection, filter interface{}) (*mongo.Cursor, error) {

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	return cursor, nil
}
