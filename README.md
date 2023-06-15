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
    url := mongodb.LoadEnv("mongourl") // mongourl is the env variable name

	db := mongodb.Connect(url)
}

```

