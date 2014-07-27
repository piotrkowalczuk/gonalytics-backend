package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	MIN_VISIT_DURATION = 30 * time.Minute
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

func VisitsAverageDuration(visits []*Visit) time.Duration {
	var averageDuration float64
	var overallDuration int64

	if len(visits) == 0 {
		return 0
	}

	for _, visit := range visits {
		overallDuration += visit.LastActionAt.Sub(visit.FirstActionAt).Nanoseconds()
	}

	averageDuration = float64(overallDuration) / float64(len(visits))
	return time.Duration(averageDuration)
}

func VisitsGroupedByFirstActionAt(visits []*Visit) []*AmountInTime {
	dateFormat := "2006-01-02 15"
	groupedVisits := make(map[string]int64)
	visitsNumber := []*AmountInTime{}

	for _, visit := range visits {
		dateString := visit.FirstActionAt.UTC().Format(dateFormat)
		if _,ok := groupedVisits[dateString]; ok {
		    groupedVisits[dateString] += 1
		} else {
			groupedVisits[dateString] = 1
		}
	}

	for dateString, nbOfVisits := range groupedVisits {
		dateTime, _ := time.Parse(dateFormat, dateString)
		visitNumber := AmountInTime{
			Amount: nbOfVisits,
			DateTime: dateTime,
		}

		visitsNumber = append(visitsNumber, &visitNumber)
	}

	return visitsNumber
}

func VisitsGroupedLocationCountryCode(visits []*Visit) map[string]int64 {
	grouped := make(map[string]int64)

	for _, visit := range visits {
		if _,ok := grouped[visit.Location.CountryCode]; ok {
			grouped[visit.Location.CountryCode] += 1
		} else {
			grouped[visit.Location.CountryCode] = 1
		}
	}

	return grouped
}
