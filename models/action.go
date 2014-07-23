package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Action struct {
	Id              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Referrer        string        `json:"referrer" bson:"referrer"`
	Page            *Page         `json:"page" bson:"page"`
	CreatedAt       time.Time     `json:"createdAt" bson:"created_at"`
	CreatedAtBucket []string      `json:"createdAtBucket" bson:"created_at_bucket"`
}
