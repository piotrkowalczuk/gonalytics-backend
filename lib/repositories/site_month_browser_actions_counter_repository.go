package repositories

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

const (
	// SiteMonthBrowserActionsCounterColumnFamily ...
	SiteMonthBrowserActionsCounterColumnFamily = "site_month_browser_actions_counter"
	// SiteMonthBrowserActionsFields ...
	SiteMonthBrowserActionsFields = `
        site_id, count, browser_name, browser_version,
        made_at_year, made_at_month
    `
)

// SiteMonthBrowserActionsCounterRepository ...
type SiteMonthBrowserActionsCounterRepository struct {
	Repository
}

// Increment ...
func (smbacr *SiteMonthBrowserActionsCounterRepository) Increment(
	siteID int64,
	browserName string,
	browserVersion string,
	date time.Time,
) error {
	cql := `
    UPDATE ` + SiteMonthBrowserActionsCounterColumnFamily + `
    SET count = count + 1
    WHERE site_id = ?
    AND browser_name = ?
    AND browser_version = ?
    AND made_at_year = ?
    AND made_at_month = ?
    `
	return smbacr.Repository.
		Cassandra.
		Query(cql, siteID, browserName, browserVersion, date.Year(), date.Month()).
		Exec()
}

// Find ...
func (smbacr *SiteMonthBrowserActionsCounterRepository) Find(
	siteID int64,
	date time.Time,
) ([]*models.SiteMonthBrowserActionsCounterEntity, error) {
	cql := `SELECT ` + SiteMonthBrowserActionsFields +
		` FROM ` + SiteMonthBrowserActionsCounterColumnFamily +
		` WHERE site_id = ? AND made_at_year = ? AND made_at_month = ?`

	iter := smbacr.Repository.
		Cassandra.
		Query(cql, siteID, date.Year(), date.Month()).
		Consistency(gocql.One).
		Iter()

	var counter models.SiteMonthBrowserActionsCounterEntity
	counters := []*models.SiteMonthBrowserActionsCounterEntity{}

	for iter.Scan(
		&counter.SiteID,
		&counter.Count,
		&counter.BrowserName,
		&counter.BrowserVersion,
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
