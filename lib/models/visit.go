package models

import (
	"time"

	"github.com/gocql/gocql"
)

const (
	// MinVisitDuration represents minimal duration of single visit
	MinVisitDuration = 30 * time.Minute
)

// Visit ...
type Visit struct {
	ID          gocql.UUID `json:"id,omitempty" cql:"id"`
	IP          string     `json:"ip,omitempty" cql:"ip"`
	NbOfActions int64      `json:"nbOfActions,omitempty" cql:"nb_of_actions"`
	SiteID      int64      `json:"siteId,omitempty" cql:"site_id"`
	Referrer    string     `json:"referrer,omitempty" cql:"referrer"`
	Language    string     `json:"language,omitempty" cql:"language"`
	// Browser         *Browser         `json:"browser,omitempty" cql:"browser"`
	// Screen          *Screen          `json:"screen,omitempty" cql:"screen"`
	// OperatingSystem *OperatingSystem `json:"os,omitempty" cql:"os"`
	// Device          *Device          `json:"device,omitempty" cql:"device"`
	// Location        *Location        `json:"location,omitempty" cql:"location"`
	// FirstPage       *Page            `json:"firstPage,omitempty" cql:"first_page"`
	// LastPage        *Page            `json:"lastPage,omitempty" cql:"last_page"`
	FirstActionAt time.Time `json:"firstActionAt,omitempty" cql:"first_action_at"`
	LastActionAt  time.Time `json:"lastActionAt,omitempty" cql:"last_action_at"`
	// Fields not related to database, comes from JOIN's etc.
	// Actions Actions `json:"actions,omitempty" cql:"-"`
}

// Visits ...
type Visits []*Visit

// Length ...
func (v *Visits) Length() int {
	return len(*v)
}

// GetIDs ...
func (v *Visits) GetIDs() (IDs []*gocql.UUID) {
	for _, visit := range *v {
		IDs = append(IDs, &visit.ID)
	}

	return
}

// GetByID returns first Visit object with given ID.
func (v *Visits) GetByID(ID gocql.UUID) (*Visit, error) {
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
		overallDuration += visit.LastActionAt.Sub(visit.FirstActionAt).Nanoseconds()
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

	// for _, visit := range *v {
	// 	if _, ok := distribution[visit.Location.CountryCode]; ok {
	// 		distribution[visit.Location.CountryCode]++
	// 	} else {
	// 		distribution[visit.Location.CountryCode] = 1
	// 	}
	// }

	for countryCode, nbOfVisits := range distribution {
		visitNumber := &AmountInCountry{
			Amount:      nbOfVisits,
			CountryCode: countryCode,
		}

		amount = append(amount, visitNumber)
	}

	return
}
