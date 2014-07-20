package models

import (
	"labix.org/v2/mgo/bson"
)

type Visit struct {
	Id         bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	ActionsIds []bson.ObjectId `json:"actions" bson:"actions"`
	CreatedAt  *MongoDate      `json:"createdAt" bson:"created_at"`
}
