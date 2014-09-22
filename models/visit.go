package models

import (
	"time"

	"labix.org/v2/mgo/bson"
)

const (
	// MinVisitDuration represents minimal duration of single visit
	MinVisitDuration = 30 * time.Minute
)

// Visit ...
type Visit struct {
	ID                  bson.ObjectId    `json:"id,omitempty" bson:"_id,omitempty"`
	IP                  string           `json:"ip,omitempty" bson:"ip,omitempty"`
	NbOfActions         int64            `json:"nbOfActions,omitempty" bson:"nb_of_actions"`
	SiteID              int64            `json:"siteId,omitempty" bson:"site_id"`
	Referrer            string           `json:"referrer,omitempty" bson:"referrer"`
	Language            string           `json:"language,omitempty" bson:"language"`
	Browser             *Browser         `json:"browser,omitempty" bson:"browser"`
	Screen              *Screen          `json:"screen,omitempty" bson:"screen"`
	OperatingSystem     *OperatingSystem `json:"os,omitempty" bson:"os"`
	Device              *Device          `json:"device,omitempty" bson:"device"`
	Location            *Location        `json:"location,omitempty" bson:"location"`
	FirstPage           *Page            `json:"firstPage,omitempty" bson:"first_page"`
	LastPage            *Page            `json:"lastPage,omitempty" bson:"last_page"`
	FirstActionAt       *time.Time       `json:"firstActionAt,omitempty" bson:"first_action_at"`
	FirstActionAtBucket []string         `json:"firstActionAtBucket,omitempty" bson:"first_action_at_bucket"`
	LastActionAt        *time.Time       `json:"lastActionAt,omitempty" bson:"last_action_at"`
	LastActionAtBucket  []string         `json:"lastActionAtBucket,omitempty" bson:"last_action_at_bucket"`
	// Fields not related to database, comes from JOIN's etc.
	Actions Actions `json:"actions,omitempty" bson:"-"`
}

// Visits ...
type Visits []*Visit

// Length ...
func (v *Visits) Length() int {
	return len(*v)
}

// GetIDs ...
func (v *Visits) GetIDs() (IDs []*bson.ObjectId) {
	for _, visit := range *v {
		IDs = append(IDs, &visit.ID)
	}

	return
}

// GetByID returns first Visit object with given ID.
func (v *Visits) GetByID(ID bson.ObjectId) (*Visit, error) {
	var err error

	for _, visit := range *v {
		if visit.ID == ID {
			return visit, nil
		}
	}

	return nil, err
}

// VisitsAverageDuration ...
func (v *Visits) VisitsAverageDuration() time.Duration {
	var averageDuration float64
	var overallDuration int64

	if v.Length() == 0 {
		return 0
	}

	for _, visit := range *v {
		overallDuration += visit.LastActionAt.Sub(*visit.FirstActionAt).Nanoseconds()
	}

	averageDuration = float64(overallDuration) / float64(v.Length())
	return time.Duration(averageDuration)
}

// MapToDistributionByTime ...
func (v *Visits) MapToDistributionByTime() []*AmountInTime {
	dateFormat := "2006-01-02 15"
	distribution := make(map[string]int64)
	visitsNumber := []*AmountInTime{}

	for _, visit := range *v {
		dateString := visit.FirstActionAt.UTC().Format(dateFormat)
		if _, ok := distribution[dateString]; ok {
			distribution[dateString]++
		} else {
			distribution[dateString] = 1
		}
	}

	for dateString, nbOfVisits := range distribution {
		dateTime, _ := time.Parse(dateFormat, dateString)
		visitNumber := AmountInTime{
			Amount:   nbOfVisits,
			DateTime: dateTime,
		}

		visitsNumber = append(visitsNumber, &visitNumber)
	}

	return visitsNumber
}

// MapToDistributionByCountryCode ...
func (v *Visits) MapToDistributionByCountryCode() (amount []*AmountInCountry) {
	distribution := make(map[string]int64)

	for _, visit := range *v {
		if _, ok := distribution[visit.Location.CountryCode]; ok {
			distribution[visit.Location.CountryCode]++
		} else {
			distribution[visit.Location.CountryCode] = 1
		}
	}

	for countryCode, nbOfVisits := range distribution {
		visitNumber := &AmountInCountry{
			Amount:      nbOfVisits,
			CountryCode: countryCode,
		}

		amount = append(amount, visitNumber)
	}

	return
}
