package repositories

import (
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// VisitCollection ...
const VisitCollection = "visit"

// VisitRepository ...
type VisitRepository struct {
	Repository
}

// Collection ...
func (vr *VisitRepository) Collection() *mgo.Collection {
	return vr.Repository.MongoDB.C(VisitCollection)
}

// Find ...
func (vr *VisitRepository) Find(query bson.M) *mgo.Query {
	return vr.Collection().Find(query)
}

// Count returns number of visit for given date range
func (vr *VisitRepository) Count(dateTimeRange string) (int64, error) {
	nbOfVisits, err := vr.Collection().
		Find(bson.M{"last_action_at_bucket": dateTimeRange}).
		Count()

	return int64(nbOfVisits), err
}

// CountByCountryID returns number of countries that user visit from
// for given date range
func (vr *VisitRepository) CountByCountryID(dateTimeRange string) (int64, error) {
	var result = struct {
		NbOfCountries int64 `bson:"nb_of_countries"`
	}{NbOfCountries: 0}

	pipeline := []bson.M{
		{"$match": bson.M{"last_action_at_bucket": dateTimeRange}},
		{"$group": bson.M{"_id": "$location.country_id"}},
		{"$group": bson.M{
			"_id":             0,
			"nb_of_countries": bson.M{"$sum": 1}}},
		{"$project": bson.M{
			"_id":             0,
			"nb_of_countries": 1,
		}},
	}

	iter := vr.Collection().Pipe(pipeline).Iter()
	iter.Next(&result)

	if iter.Err() != nil {
		return 0, iter.Err()
	}
	return result.NbOfCountries, nil
}

// DistributionByTime returns number of visits grouped by hours
func (vr *VisitRepository) DistributionByTime(dateTimeRange string) ([]*models.AmountInTime, error) {
	var visits models.Visits

	err := vr.Collection().Find(bson.M{
		"first_action_at_bucket": dateTimeRange},
	).Select(bson.M{
		"first_action_at": 1,
	}).All(&visits)

	if err != nil {
		return nil, err
	}

	return visits.MapToDistributionByTime(), nil
}

// DistributionByCountry returns number of visits grouped by country ID
func (vr *VisitRepository) DistributionByCountry(dateTimeRange string) ([]*models.AmountInCountry, error) {
	var visits models.Visits

	err := vr.Collection().Find(bson.M{
		"first_action_at_bucket": dateTimeRange,
	}).Select(bson.M{
		"location": 1,
	}).All(&visits)

	if err != nil {
		return nil, err
	}

	return visits.MapToDistributionByCountryCode(), nil
}
