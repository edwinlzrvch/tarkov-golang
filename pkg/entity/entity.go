package entity

import (
	"gopkg.in/mgo.v2/bson"
)

type ID bson.ObjectId

//StringToID convert a string to an ID
func StringToID(s string) ID {
	return ID(bson.ObjectIdHex(s))
}

//NewID create a new id
func NewID() ID {
	return StringToID(bson.NewObjectId().Hex())
}
