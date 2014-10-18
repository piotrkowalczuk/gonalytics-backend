package repositories

import "github.com/gocql/gocql"

// Repository ...
type Repository struct {
	Cassandra *gocql.Session
}
