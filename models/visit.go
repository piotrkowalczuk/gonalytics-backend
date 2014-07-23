package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Visit struct {
	Id                  bson.ObjectId    `json:"id" bson:"_id,omitempty"`
	Actions             []*Action        `json:"actions" bson:"actions"`
	NbOfActions         int64            `json:"nbOfActions" bson:"nb_of_actions"`
	SiteId              int64            `json:"siteId" bson:"site_id"`
	Referrer            string           `json:"referrer" bson:"referrer"`
	Language            string           `json:"language" bson:"language"`
	Browser             *Browser         `json:"browser" bson:"browser"`
	Screen              *Screen          `json:"screen" bson:"screen"`
	OperatingSystem     *OperatingSystem `json:"os" bson:"os"`
	Device              *Device          `json:"device" bson:"device"`
	Location            *Location        `json:"location" bson:"location"`
	CreatedAt           time.Time        `json:"createdAt" bson:"created_at"`
	CreatedAtBucket     []string         `json:"createdAtBucket" bson:"created_at_bucket"`
	FirstPage           *Page            `json:"firstPage" bson:"first_page"`
	LastPage            *Page            `json:"lastPage" bson:"last_page"`
	FirstActionAt       time.Time        `json:"firstActionAt" bson:"first_action_at"`
	FirstActionAtBucket []string         `json:"firstActionAtBucket" bson:"first_action_at_bucket"`
	LastActionAt        time.Time        `json:"lastActionAt" bson:"last_action_at"`
	LastActionAtBucket  []string         `json:"lastActionAtBucket" bson:"last_action_at_bucket"`
}
