package entity

type Room struct {
	Id          ID     `json:"id" bson:"_id,omitempty"`
	Rate        int    `json:"rate" bson:"rate"`
	Host        string `json:"host" bson:"host"`
	Description string `json:"description" bson:"description"`
	PLayers     int32  `json:"players" bson:"players"`
}
