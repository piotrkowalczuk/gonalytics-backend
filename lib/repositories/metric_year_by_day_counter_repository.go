package repositories

import (
	"time"

	"github.com/gocql/gocql"
)

const (
	// MetricYearByDayCounterFields ...
	MetricYearByDayCounterFields = `
    dimensions_names, dimensions_values, made_at_day,
	made_at_year, count,
	`
)

// MetricYearByDayCounterRepository ...
type MetricYearByDayCounterRepository struct {
	Repository
	ColumnFamily string
}

// NewMetricYearByDayCounterRepository ...
func NewMetricYearByDayCounterRepository(repository Repository, columnFamily string) MetricYearByDayCounterRepository {
	return MetricYearByDayCounterRepository{
		Repository:   repository,
		ColumnFamily: columnFamily,
	}
}

// IncrementQuery ...
func (mybdcr *MetricYearByDayCounterRepository) IncrementQuery(dimensionsNames, dimensionsValues string, now time.Time) *gocql.Query {
	cql := `
    UPDATE ` + mybdcr.ColumnFamily + `
    SET count = count + 1
	WHERE dimensions_names = ?
    AND dimensions_values = ?
    AND made_at_year = ?
    AND made_at_day = ?
    `
	from := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)
	nbOfDays := int(now.Sub(from).Hours()) / 24

	return mybdcr.Repository.
		Cassandra.
		Query(cql, dimensionsNames, dimensionsValues, now.Year(), nbOfDays)
}

// Increment ...
func (mybdcr *MetricYearByDayCounterRepository) Increment(dimensionsNames, dimensionsValues string, date time.Time) error {
	return mybdcr.IncrementQuery(dimensionsNames, dimensionsValues, date).Exec()
}
