package mongo

import (
	"github.com/globalsign/mgo/bson"
)

//IDableObjectID is a type that has DbID which returns the id field for mongo
type IDableObjectID interface {
	DbID() bson.ObjectId
}

//IDableString is a type that has DbID which returns the id field for mongo
type IDableString interface {
	DbID() string
}
