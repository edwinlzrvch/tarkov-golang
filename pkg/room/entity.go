package room

import "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"

type Room struct {
	Id          string    `json:"id" bson:"_id,omitempty"`
	Rate        int       `json:"rate" bson:"rate"`
	Host        string    `json:"host" bson:"host"`
	Description string    `json:"description" bson:"description"`
	PLayers     int       `json:"players" bson:"players"`
	Uid         uuid.UUID `json:"uid" bson:"uid"`
}
