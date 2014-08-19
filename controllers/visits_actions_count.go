package controllers

import (
	"net/http"

	"labix.org/v2/mgo/bson"
)

// VisitsActionsCountController ...
type VisitsActionsCountController struct {
	GeneralController
	ResponseData struct {
		NbOfActions int64 `json:"nbOfActions" bson:"nb_of_actions"`
	}
}

// Get ...
func (vacc *VisitsActionsCountController) Get() {
	dateTimeRange := vacc.GetString("dateTimeRange")
	pipeline := []bson.M{}

	if dateTimeRange != "" {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"actions.created_at_bucket": dateTimeRange}})
	}

	pipeline = append(pipeline, bson.M{"$group": bson.M{
		"_id":           bson.M{},
		"nb_of_actions": bson.M{"$sum": "$nb_of_actions"},
	}})
	pipeline = append(pipeline, bson.M{"$project": bson.M{
		"_id":           0,
		"nb_of_actions": 1,
	}})

	iter := vacc.MongoPool.Collection("visit").Pipe(pipeline).Iter()
	iter.Next(&vacc.ResponseData)

	vacc.AbortIf(iter.Err(), "Unexpected error.", http.StatusInternalServerError)
}
