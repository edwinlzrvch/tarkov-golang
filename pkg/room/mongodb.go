package room

import (
	. "../entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type repo struct {
	db  *mongo.Database
	ctx context.Context
}

func NewMongoRepository(db *mongo.Database, ctx context.Context) Repository {
	return &repo{
		db,
		ctx,
	}
}

func (r *repo) Find(id string) (*Room, error) {
	result := Room{}
	coll := r.db.Collection("Rooms")
	err := coll.FindOne(r.ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return &result, nil
}

func (r *repo) FindAll() []*Room {
	var allRooms []*Room
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

func (r *repo) Add(room *Room) (*mongo.InsertOneResult, error) {
	coll := r.db.Collection("Rooms")
	roomResult, err := coll.InsertOne(r.ctx, room)
	println(roomResult.InsertedID)
	if err != nil {
		return nil, err
	}
	return roomResult, nil
}
