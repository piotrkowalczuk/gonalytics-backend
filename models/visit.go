package models

import (
	"labix.org/v2/mgo/bson"
)

type Visit struct {
	Id              bson.ObjectId    `json:"id" bson:"_id,omitempty"`
	Actions         []*Action        `json:"actions" bson:"actions"`
	NbOfActions     int64            `json:"nbOfActions" bson:"nb_of_actions"`
	SiteId          int64            `json:"siteId" bson:"site_id"`
	Referrer        string           `json:"referrer" bson:"referrer"`
	Language        string           `json:"language" bson:"language"`
	Browser         *Browser         `json:"browser" bson:"browser"`
	Screen          *Screen          `json:"screen" bson:"screen"`
	OperatingSystem *OperatingSystem `json:"os" bson:"os"`
	Device          *Device          `json:"device" bson:"device"`
	Location        *Location        `json:"location" bson:"location"`
	CreatedAt       *MongoDate       `json:"createdAt" bson:"created_at"`
	FirstPage       *Page            `json:"firstPage" bson:"first_page"`
	LastPage        *Page            `json:"lastPage" bson:"last_page"`
	FirstActionAt   *MongoDate       `json:"firstActionAt" bson:"first_action_at"`
	LastActionAt    *MongoDate       `json:"lastActionAt" bson:"last_action_at"`
}
