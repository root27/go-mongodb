package mongodb

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func TestFindOne(t *testing.T) {

	client, err := Connect("URI")

	if err != nil {
		t.Errorf("Error connecting to MongoDB: %v", err)
	}

	collection := client.GetCollection("MongoTest", "user")

	if collection == nil {
		t.Errorf("Collection is nil")
	}

	filter := map[string]interface{}{
		"name": "Oguzhan",
	}

	result := collection.FindOne(filter)

	if result == nil {
		t.Errorf("SingleResult is nil")
	}

	if result.Err() != nil {
		t.Errorf("Error in SingleResult: %v", result.Err())
	}

	var UserRes bson.M

	err = result.Decode(&UserRes)

	if err != nil {
		t.Errorf("Error decoding result: %v", err)
	}

	if UserRes == nil {
		t.Errorf("Result is nil")
	}

	if UserRes["name"] != "Oguzhan" {
		t.Errorf("Result is not equal to expected result")
	}

}
