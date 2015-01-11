package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteDayBrowserActionsCounterColumnFamily ...
	SiteDayBrowserActionsCounterColumnFamily = "site_day_browser_actions_counter"
	// SiteDayBrowserActionsFields ...
	SiteDayBrowserActionsFields = `
        site_id, count, browser_name, browser_version,
        made_at_year, made_at_month, made_at_day
    `
)

// SiteDayBrowserActionsCounterRepository ...
type SiteDayBrowserActionsCounterRepository struct {
	Repository
}

// Increment ...
func (sdbacr *SiteDayBrowserActionsCounterRepository) Increment(
	siteID int64,
	browserName string,
	browserVersion string,
	date time.Time,
) error {
	cql := `
    UPDATE ` + SiteDayBrowserActionsCounterColumnFamily + `
    SET count = count + 1
    WHERE site_id = ?
    AND browser_name = ?
    AND browser_version = ?
    AND made_at_year = ?
    AND made_at_month = ?
    AND made_at_day = ?
    `
	return sdbacr.Repository.
		Cassandra.
		Query(cql, siteID, browserName, browserVersion, date.Year(), date.Month(), date.Day()).
		Exec()
}

// Find ...
func (sdbacr *SiteDayBrowserActionsCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteDayBrowserActionsCounterEntity, error) {
	cql := `SELECT ` + SiteDayBrowserActionsFields +
		` FROM ` + SiteDayBrowserActionsCounterColumnFamily +
		` WHERE site_id = ? AND made_at_year = ? AND made_at_month = ? AND made_at_day = ?`

	iter := sdbacr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year(), date.Month(), date.Day()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteDayBrowserActionsCounterEntity
	counters := []*models.SiteDayBrowserActionsCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.Count,
		&counter.BrowserName,
		&counter.BrowserVersion,
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
