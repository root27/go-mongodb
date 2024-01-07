package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (c Collection) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {

	result, err := c.collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c Collection) UpdateMany(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {

	result, err := c.collection.UpdateMany(context.TODO(), filter, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}
