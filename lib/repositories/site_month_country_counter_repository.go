package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteMonthCountryActionsCounterColumnFamily ...
	SiteMonthCountryActionsCounterColumnFamily = "site_month_country_actions_counter"
	// SiteMonthCountryVisitsCounterColumnFamily ...
	SiteMonthCountryVisitsCounterColumnFamily = "site_month_country_visits_counter"
	// SiteMonthCountryCounterFields ...
	SiteMonthCountryCounterFields = `
        site_id, count, location_country_name, location_country_code, location_country_id,
        made_at_year, made_at_month
    `
)

// SiteMonthCountryCounterRepository ...
type SiteMonthCountryCounterRepository struct {
	Repository
}

// Increment ...
func (smccr *SiteMonthCountryCounterRepository) Increment(siteID int64, countryName string, countryCode string, countryID uint, date time.Time) error {

	cql := `
    UPDATE ` + smccr.ColumnFamily + `
    SET count = count + 1
    WHERE site_id = ?
    AND location_country_name = ?
    AND location_country_code = ?
    AND location_country_id = ?
    AND made_at_year = ?
    AND made_at_month = ?
    `
	return smccr.Repository.
		Cassandra.
		Query(cql, siteID, countryName, countryCode, countryID, date.Year(), date.Month()).
		Exec()
}

// Find ...
func (smccr *SiteMonthCountryCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteMonthCountryCounterEntity, error) {
	cql := `SELECT ` + SiteMonthCountryCounterFields +
		` FROM ` + smccr.ColumnFamily +
		` WHERE site_id = ? AND made_at_year = ? AND made_at_month = ?`

	iter := smccr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year(), date.Month()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteMonthCountryCounterEntity
	counters := []*models.SiteMonthCountryCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.Count,
		&counter.LocationCountryName,
		&counter.LocationCountryCode,
		&counter.LocationCountryID,
		&counter.MadeAtYear,
		&counter.MadeAtMonth,
	) {
		counters = append(counters, &counter)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}

	return counters, nil
}
