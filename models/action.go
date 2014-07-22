package models

import (
	"labix.org/v2/mgo/bson"
)

type Action struct {
	Id        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Referrer  string        `json:"referrer" bson:"referrer"`
	Page      *Page         `json:"page" bson:"page"`
	CreatedAt *MongoDate    `json:"createdAt" bson:"created_at"`
}
