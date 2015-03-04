package repositories

import (
	"time"

	"github.com/gocql/gocql"
)

const (
	// MetricDayCounterFields ...
	MetricDayCounterFields = `
    dimensions_names, dimensions_values, made_at_day, made_at_month,
    made_at_year, count,
	`
)

// MetricDayCounterRepository ...
type MetricDayCounterRepository struct {
	Repository
	ColumnFamily string
}

// IncrementQuery ...
func (mdcr *MetricDayCounterRepository) IncrementQuery(dimensionsNames, dimensionsValues string, date time.Time) *gocql.Query {
	cql := `
    UPDATE ` + mdcr.ColumnFamily + `
    SET count = count + 1
	WHERE dimensions_names = ?
    AND dimensions_values = ?
    AND made_at_year = ?
    AND made_at_month = ?
    AND made_at_day = ?
    `

	return mdcr.Repository.
		Cassandra.
		Query(cql, dimensionsNames, dimensionsValues, date.Year(), date.Month(), date.Day())
}

// Increment ...
func (mdcr *MetricDayCounterRepository) Increment(dimensionsNames, dimensionsValues string, date time.Time) error {
	return mdcr.IncrementQuery(dimensionsNames, dimensionsValues, date).Exec()
}
