package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateOne(collection mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateMany(collection mongo.Collection, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {

	result, err := collection.UpdateMany(context.TODO(), filter, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}
