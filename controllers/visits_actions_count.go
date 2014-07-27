package controllers

import (
	"labix.org/v2/mgo/bson"
	"net/http"
)

type VisitsActionsCountController struct {
	BaseController
}

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

	var result struct {
		NbOfActions int64 `bson:"nb_of_actions"`
	}

	iter := vacc.MongoPool.Collection("visit").Pipe(pipeline).Iter()
	iter.Next(&result)
	err := iter.Err()

	vacc.abortIf(err, http.StatusInternalServerError)
	vacc.Data["json"] = result.NbOfActions
	vacc.ServeJson()
}
