package repositories

import (
	"time"

	"github.com/gocql/gocql"
)

// Repository ...
type Repository struct {
	Cassandra *gocql.Session
}

// MetricIncrementer ...
type MetricIncrementer interface {
	Increment(string, string, time.Time) error
}
