package controllers

import (
	"labix.org/v2/mgo/bson"
	"net/http"
)

type VisitsCountriesCountController struct {
	BaseController
}

func (vccc *VisitsCountriesCountController) Get() {
	dateTimeRange := vccc.GetString("dateTimeRange")
	pipeline := []bson.M{
		{"$match": bson.M{"actions.created_at_bucket": dateTimeRange}},
		{"$group": bson.M{"_id": "$location.country_id"}},
		{"$group": bson.M{
			"_id":             0,
			"nb_of_countries": bson.M{"$sum": 1}}},
		{"$project": bson.M{
			"_id":             0,
			"nb_of_countries": 1,
		}},
	}

	var result struct {
		NbOfCountries int64 `bson:"nb_of_countries"`
	}

	iter := vccc.MongoPool.Collection("visit").Pipe(pipeline).Iter()
	iter.Next(&result)
	err := iter.Err()

	vccc.abortIf(err, http.StatusInternalServerError)
	vccc.Data["json"] = result.NbOfCountries
	vccc.ServeJson()
}
