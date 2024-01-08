package mongodb

import (
	"testing"
)

func TestConnect(t *testing.T) {

	client, err := Connect("URI")

	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}

	if client == nil {
		t.Errorf("Client is nil")
	}

}

func TestGetCollection(t *testing.T) {

	client, err := Connect("URI")

	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}

	collection := client.GetCollection("MongoTest", "user")

	if collection == nil {
		t.Errorf("Collection is nil")
	}

}
