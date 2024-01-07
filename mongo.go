package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	client *mongo.Client
}

type Collection struct {
	collection *mongo.Collection
}

func (c Client) GetCollection(dbName string, collectionName string) Collection {

	collection := c.client.Database(dbName).Collection(collectionName)

	newCollection := Collection{collection: collection}

	return newCollection
}

func (c Collection) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {

	result, err := c.collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	return result, nil
}
