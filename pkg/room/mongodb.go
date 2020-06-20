package room

import (
	"github.com/juju/mgosession"
	"gopkg.in/mgo.v2/bson"
)

type repo struct {
	pool *mgosession.Pool
}

func NewMongoRepository(p *mgosession.Pool) Repository {
	return &repo{
		pool: p,
	}
}

func (r *repo) Find(id string) (*Room, error) {
	result := Room{}
	session := r.pool.Session(nil)
	coll := session.DB("").C("Rooms")
	err := coll.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repo) Add(room *Room) (*Room, error) {
	session := r.pool.Session(nil)
	coll := session.DB("TarkvoDb").C("Rooms")
	err := coll.Insert(room)
	if err != nil {
		return nil, err
	}
	return room, nil
}
