## Go MongoDB

---

### Description

This is a simple package to connect to a MongoDB database.

### Usage

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

