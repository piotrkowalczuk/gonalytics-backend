package controllers

import (
	"labix.org/v2/mgo/bson"
)

type VisistsActionsCountController struct {
	BaseController
}

func (vacc *VisistsActionsCountController) Get() {
	dateTimeRange := vacc.GetString("dateTimeRange")
	pipeline := []bson.M{
		{"$match": bson.M{"actions.created_at.bucket": dateTimeRange}},
		{"$group": bson.M{
			"_id":           bson.M{},
			"nb_of_actions": bson.M{"$sum": "$nb_of_actions"},
		}},
		{"$project": bson.M{
			"_id":           0,
			"nb_of_actions": 1,
		}},
	}

	var result struct {
		NbOfActions int64 `bson:"nb_of_actions"`
	}

	iter := vacc.MongoPool.Collection("visit").Pipe(pipeline).Iter()
	iter.Next(&result)
	err := iter.Err()

	if err != nil {
		vacc.Abort("500")
	}

	vacc.Data["json"] = result.NbOfActions
	vacc.ServeJson()
}
