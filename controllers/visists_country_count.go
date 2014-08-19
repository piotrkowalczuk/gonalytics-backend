package controllers

import (
	"net/http"

	"labix.org/v2/mgo/bson"
)

// VisitsCountriesCountController ...
type VisitsCountriesCountController struct {
	GeneralController
	ResponseData struct {
		NbOfCountries int64 `bson:"nb_of_countries"`
	}
}

// Get ...
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

	iter := vccc.MongoPool.Collection("visit").Pipe(pipeline).Iter()
	iter.Next(&vccc.ResponseData)

	vccc.AbortIf(iter.Err(), "Unexpected error.", http.StatusInternalServerError)
}
