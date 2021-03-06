package entity

// Room interface
type Room struct {
	ID          ID     `json:"id" bson:"_id,omitempty"`
	Rate        int32  `json:"rate" bson:"rate"`
	Host        string `json:"host" bson:"host"`
	Description string `json:"description" bson:"description"`
	Players     int32  `json:"players" bson:"players"`
}
