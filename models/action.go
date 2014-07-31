package models

import (
	"time"

	"labix.org/v2/mgo/bson"
)

// Action ...
type Action struct {
	ID              bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Referrer        string        `json:"referrer" bson:"referrer"`
	Page            *Page         `json:"page" bson:"page"`
	CreatedAt       *time.Time    `json:"createdAt" bson:"created_at"`
	CreatedAtBucket []string      `json:"createdAtBucket" bson:"created_at_bucket"`
}
