package repositories

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// ActionCollection ...
const ActionCollection = "action"

// ActionRepository ...
type ActionRepository struct {
	Repository
}

// Collection ...
func (ar *ActionRepository) Collection() *mgo.Collection {
	return ar.Repository.MongoDB.C(ActionCollection)
}

// Find ...
func (ar *ActionRepository) Find(query bson.M) *mgo.Query {
	return ar.Collection().Find(query)
}
