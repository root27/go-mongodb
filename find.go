package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func (c *Collection) FindOne(filter interface{}) *mongo.SingleResult {

	result := c.collection.FindOne(context.TODO(), filter)

	return result
}

func (c *Collection) FindMany(filter interface{}) (*mongo.Cursor, error) {

	cursor, err := c.collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	return cursor, nil
}
