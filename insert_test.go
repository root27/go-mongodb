package mongodb

import (
	"testing"
)

func TestInsert(t *testing.T) {

	client, err := Connect("URI")

	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}

	collection := client.GetCollection("MongoTest", "user")

	if collection == nil {
		t.Errorf("Collection is nil")
	}

	data := map[string]interface{}{
		"name": "test",
		"age":  20,
	}

	result, err := collection.InsertOne(data)

	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}

	if result == nil {
		t.Errorf("InsertOneResult is nil")
	}

	if result.InsertedID == nil {
		t.Errorf("InsertedID is nil")
	}

}

func TestInsertMany(t *testing.T) {

	client, err := Connect("URI")

	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}

	collection := client.GetCollection("MongoTest", "user")

	if collection == nil {
		t.Errorf("Collection is nil")
	}

	data := []interface{}{
		map[string]interface{}{
			"name": "test",
			"age":  10,
		},
		map[string]interface{}{
			"name": "test2",
			"age":  30,
		},
	}

	result, err := collection.InsertMany(data)

	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}

	if result == nil {
		t.Errorf("InsertManyResult is nil")
	}

	if result.InsertedIDs == nil {
		t.Errorf("InsertedIDs is nil")
	}

	if len(result.InsertedIDs) != 2 {
		t.Errorf("InsertedIDs length is not 2")
	}

}
