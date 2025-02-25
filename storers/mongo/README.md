# REST Layer MongoDB Backend

[![Go Reference](https://pkg.go.dev/badge/github.com/searis/rest-layer/storers/mongo.svg)](https://pkg.go.dev/github.com/searis/rest-layer/storers/mongo)

This [REST Layer](https://github.com/searis/rest-layer) resource storage backend stores data in a MongoDB cluster using [mgo](https://godoc.org/labix.org/v2/mgo).

## Usage

```go
import "github.com/rs/rest-layer-mongo"
```

Create a mgo master session:

```go
session, err := mgo.Dial(url)
```

Create a resource storage handler with a given DB/collection:

```go
s := mongo.NewHandler(session, "the_db", "the_collection")
```

Use this handler with a resource:

```go
index.Bind("foo", foo, s, resource.DefaultConf)
```

You may want to create a many mongo handlers as you have resources as long as you want each resources in a different collection. You can share the same `mgo` session across all you handlers.

### Object ID

This package also provides a REST Layer [schema.Validator](https://godoc.org/github.com/searis/rest-layer/schema#Validator) for MongoDB ObjectIDs. This validator ensures proper binary serialization of the Object ID in the database for space efficiency.

You may reference this validator using [mongo.ObjectID](https://godoc.org/github.com/rs/rest-layer-mongo#ObjectID) as [schema.Field](https://godoc.org/github.com/searis/rest-layer/schema#Field).

A `mongo.NewObjectID` field hook and `mongo.ObjectIDField` helper are also provided.
