package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {

	collection := client.Database(dbName).Collection(collectionName)

	return collection
}

func DeleteOne(collection *mongo.Collection, filter interface{}) (*mongo.DeleteResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return result, nil
}
