package room

import (
	"context"
	"log"

	"github.com/edwinlzrvch/tarkov-golang/pkg/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	db  *mongo.Database
	ctx context.Context
}

// NewMongoRepository asd
func NewMongoRepository(ctx context.Context, db *mongo.Database) Repository {
	return &repo{
		db,
		ctx,
	}
}

func (r *repo) Find(id string) (*entity.Room, error) {
	result := entity.Room{}
	coll := r.db.Collection("Rooms")
	err := coll.FindOne(r.ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return &result, nil
}

// FindAll dadad
func (r *repo) FindAll() []*entity.Room {
	var allRooms []*entity.Room
	coll := r.db.Collection("Rooms")
	cursor, err := coll.Find(r.ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(r.ctx, &allRooms); err != nil {
		log.Fatal(err)
	}
	return allRooms
}

// Add daad
func (r *repo) Add(room *entity.Room) (*mongo.InsertOneResult, error) {
	coll := r.db.Collection("Rooms")
	roomResult, err := coll.InsertOne(r.ctx, room)
	println(roomResult.InsertedID)
	if err != nil {
		return nil, err
	}
	return roomResult, nil
}
