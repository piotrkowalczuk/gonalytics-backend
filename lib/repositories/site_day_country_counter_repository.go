package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteDayCountryActionsCounterColumnFamily ...
	SiteDayCountryActionsCounterColumnFamily = "site_day_country_actions_counter"
	// SiteDayCountryVisitsCounterColumnFamily ...
	SiteDayCountryVisitsCounterColumnFamily = "site_day_country_visits_counter"
	// SiteDayCountryFields ...
	SiteDayCountryCounterFields = `
        site_id, count, location_country_name, location_country_code, location_country_id,
        made_at_year, made_at_month, made_at_day
    `
)

// SiteDayCountryCounterRepository ...
type SiteDayCountryCounterRepository struct {
	Repository
}

// Increment ...
func (sdccr *SiteDayCountryCounterRepository) Increment(siteID int64, countryName string, countryCode string, countryID uint, date time.Time) error {

	cql := `
    UPDATE ` + sdccr.ColumnFamily + `
    SET count = count + 1
    WHERE site_id = ?
    AND location_country_name = ?
    AND location_country_code = ?
    AND location_country_id = ?
    AND made_at_year = ?
    AND made_at_month = ?
    AND made_at_day = ?
    `
	return sdccr.Repository.
		Cassandra.
		Query(cql, siteID, countryName, countryCode, countryID, date.Year(), date.Month(), date.Day()).
		Exec()
}

// Find ...
func (sdccr *SiteDayCountryCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteDayCountryCounterEntity, error) {
	cql := `SELECT ` + SiteDayCountryCounterFields +
		` FROM ` + sdccr.ColumnFamily +
		` WHERE site_id = ? AND made_at_year = ? AND made_at_month = ? AND made_at_day = ?`

	iter := sdccr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year(), date.Month(), date.Day()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteDayCountryCounterEntity
	counters := []*models.SiteDayCountryCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.Count,
		&counter.LocationCountryName,
		&counter.LocationCountryCode,
		&counter.LocationCountryID,
		&counter.MadeAtYear,
		&counter.MadeAtMonth,
		&counter.MadeAtDay,
	) {
		counters = append(counters, &counter)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return counters, nil
}
