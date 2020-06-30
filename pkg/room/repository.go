package room

import (
	"github.com/edwinlzrvch/tarkov-golang/pkg/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

// Reader dadada
type Reader interface {
	Find(id string) (*entity.Room, error)
	FindAll() []*entity.Room
}

// Writer datada
type Writer interface {
	Add(room *entity.Room) (*mongo.InsertOneResult, error)
	// Delete(id string) error
}

// Repository dada asd
type Repository interface {
	Reader
	Writer
}
