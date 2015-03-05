package repositories

import (
	"time"

	"github.com/gocql/gocql"
)

const (
	// MetricDayByValueCounterFields ...
	MetricDayByValueCounterFields = `
    dimensions_names, dimensions_values, made_at_day, made_at_month,
    made_at_year, count,
	`
)

// MetricDayByValueCounterRepository ...
type MetricDayByValueCounterRepository struct {
	Repository
	ColumnFamily string
}

// NewMetricDayByValueCounterRepository ...
func NewMetricDayByValueCounterRepository(repository Repository, columnFamily string) MetricDayByValueCounterRepository {
	return MetricDayByValueCounterRepository{
		Repository:   repository,
		ColumnFamily: columnFamily,
	}
}

// IncrementQuery ...
func (mdbvcr *MetricDayByValueCounterRepository) IncrementQuery(dimensionsNames, dimensionsValues string, date time.Time) *gocql.Query {
	cql := `
    UPDATE ` + mdbvcr.ColumnFamily + `
    SET count = count + 1
	WHERE dimensions_names = ?
    AND dimensions_values = ?
    AND made_at_year = ?
    AND made_at_month = ?
    AND made_at_day = ?
    `

	return mdbvcr.Repository.
		Cassandra.
		Query(cql, dimensionsNames, dimensionsValues, date.Year(), date.Month(), date.Day())
}

// Increment ...
func (mdbvcr *MetricDayByValueCounterRepository) Increment(dimensionsNames, dimensionsValues string, date time.Time) error {
	return mdbvcr.IncrementQuery(dimensionsNames, dimensionsValues, date).Exec()
}
