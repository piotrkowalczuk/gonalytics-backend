package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteYearCountryActionsCounterColumnFamily ...
	SiteYearCountryActionsCounterColumnFamily = "site_year_country_actions_counter"
	// SiteYearCountryVisitsCounterColumnFamily ...
	SiteYearCountryVisitsCounterColumnFamily = "site_year_country_visits_counter"
	// SiteYearCountryActionsFields ...
	SiteYearCountryCounterFields = `
        site_id, count, location_country_name, 
        location_country_code, location_country_id,
        made_at_year
    `
)

// SiteYearCountryCounterRepository ...
type SiteYearCountryCounterRepository struct {
	Repository
}

// Increment ...
func (syccr *SiteYearCountryCounterRepository) Increment(siteID int64, countryName string, countryCode string, countryID uint, date time.Time) error {

	cql := `
    UPDATE ` + syccr.ColumnFamily + `
    SET count = count + 1
    WHERE site_id = ?
    AND location_country_name = ?
    AND location_country_code = ?
    AND location_country_id = ?
    AND made_at_year = ?
    `
	return syccr.Repository.
		Cassandra.
		Query(cql, siteID, countryName, countryCode, countryID, date.Year()).
		Exec()
}

// Find ...
func (syccr *SiteYearCountryCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteYearCountryCounterEntity, error) {
	cql := `SELECT ` + SiteYearCountryCounterFields +
		` FROM ` + syccr.ColumnFamily +
		` WHERE site_id = ? AND made_at_year = ?`

	iter := syccr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteYearCountryCounterEntity
	counters := []*models.SiteYearCountryCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.Count,
		&counter.LocationCountryName,
		&counter.LocationCountryCode,
		&counter.LocationCountryID,
		&counter.MadeAtYear,
	) {
		counters = append(counters, &counter)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return counters, nil
}
