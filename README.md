## Go MongoDB

---

### Description

This is a simple package to connect to a MongoDB database.

### Connection and Get Collection Usage

```go
package main

import (
    "github.com/root27/go-mongodb"
)

func main(){
     

	db ,err:= mongodb.Connect("mongourl") // mongourl is the env variable name

    if err != nil {
        panic(err)
    }

    // db is a pointer to a mongo.Database
    // use it to perform operations on the database

    //Get Collection
    collection := db.GetCollection(db, "databaseName","collectionName")
}

```

### Insert Usage

```go
package main

import (
    "github.com/root27/go-mongodb"
)

func main(){
     

    db ,err:= mongodb.Connect("mongourl") // mongourl is the env variable name

    if err != nil {
        panic(err)
    }

    // db is a pointer to a mongo.Database
    // use it to perform operations on the database

    //Get Collection
    collection := db.GetCollection(db, "databaseName","collectionName")

    //Insert
    err = collection.InsertOne(db, bson.M{"name": "pi", "value": 3.14159})
    if err != nil {
        panic(err)
    }
}

```

### Find Usage

```go

package main

import (
    "github.com/root27/go-mongodb"
)

func main(){
     

    db ,err:= mongodb.Connect("mongourl") // mongourl is the env variable name

    if err != nil {
        panic(err)
    }

    // db is a pointer to a mongo.Database
    // use it to perform operations on the database

    //Get Collection
    collection := db.GetCollection(db, "databaseName","collectionName")

    //Find
    var result bson.M
    err = collection.FindOne(db, bson.M{"name": "pi"}).Decode(&result)
    if err != nil {
        panic(err)
    }

    fmt.Println("pi:", result["value"])
}

```

