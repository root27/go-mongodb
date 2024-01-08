package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (c *Collection) InsertOne(data interface{}) (*mongo.InsertOneResult, error) {

	result, err := c.collection.InsertOne(context.TODO(), data)

	return result, err
}

func (c *Collection) InsertMany(data []interface{}) (*mongo.InsertManyResult, error) {

	result, err := c.collection.InsertMany(context.TODO(), data)

	return result, err
}
