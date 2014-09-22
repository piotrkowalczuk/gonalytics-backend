package controllers

import (
	"net/http"

	"labix.org/v2/mgo/bson"
)

// ActionsCountController ...
type ActionsCountController struct {
	GeneralController
}

// ActionsCountResponse ...
type ActionsCountResponse struct {
	NbOfActions int `json:"nbOfActions" bson:"nb_of_actions"`
}

// Get ...
func (acc *ActionsCountController) Get() {
	dateTimeRange := acc.GetString("dateTimeRange")
	response := ActionsCountResponse{
		NbOfActions: 0,
	}
	pipeline := []bson.M{}

	if dateTimeRange != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"last_action_at_bucket": dateTimeRange}})
	}

	pipeline = append(pipeline, bson.M{"$group": bson.M{
		"_id":           bson.M{},
		"nb_of_actions": bson.M{"$sum": "$nb_of_actions"},
	}})
	pipeline = append(pipeline, bson.M{"$project": bson.M{
		"_id":           0,
		"nb_of_actions": 1,
	}})

	iter := acc.RepositoryManager.Visit.Collection().Pipe(pipeline).Iter()
	iter.Next(&response)

	acc.ResponseData = response
	acc.AbortIf(iter.Err(), "Unexpected error.", http.StatusInternalServerError)
}
