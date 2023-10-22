package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadEnv(envName string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(envName)
}

func Connect(envName string) (*mongo.Client, error) {

	url := LoadEnv(envName)

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client, err
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}

func Disconnect(client *mongo.Client) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func InsertOne(collection *mongo.Collection, data interface{}) (*mongo.InsertOneResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, data)
	return result, err
}

func InsertMany(collection *mongo.Collection, data []interface{}) (*mongo.InsertManyResult, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertMany(ctx, data)
	return result, err
}

func FindOne(collection *mongo.Collection, filter interface{}) *mongo.SingleResult {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result := collection.FindOne(ctx, filter)
	return result
}

func DeleteOne(collection -mongo.Collection, filter interface {}) *mongo.SingleResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result := collection.DeleteOne(ctx, filter)
	return result
}

func UpdateOne(collection -mongo.Collection, filter interface {}, update interface {}) *mongo.SingleResult {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result := collection.UpdateOne(ctx, filter, update)
	return result
}
