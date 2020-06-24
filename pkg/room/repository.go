package room

import (
	. "../entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type Reader interface {
	Find(id string) (*Room, error)
	FindAll() []*Room
}

type Writer interface {
	Add(room *Room) (*mongo.InsertOneResult, error)
	// Delete(id string) error
}

type Repository interface {
	Reader
	Writer
}
