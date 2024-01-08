package mongodb

import "testing"

func TestUpdateOne(t *testing.T) {

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

	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"age": 20,
		},
	}

	result, err := collection.UpdateOne(filter, update)

	if err != nil {
		t.Errorf("Error updating document: %v", err)
	}

	if result == nil {
		t.Errorf("UpdateResult is nil")
	}

	if result.MatchedCount != 1 {
		t.Errorf("MatchedCount is not equal to 1")
	}

	if result.ModifiedCount != 1 {
		t.Errorf("ModifiedCount is not equal to 1")
	}

	if result.UpsertedCount != 0 {
		t.Errorf("UpsertedCount is not equal to 0")
	}

	if result.UpsertedID != nil {
		t.Errorf("UpsertedID is not nil")
	}

}
