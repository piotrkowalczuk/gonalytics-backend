package repositories

import (
	"github.com/gocql/gocql"
	"labix.org/v2/mgo"
)

// Repository ...
type Repository struct {
	MongoDB   *mgo.Database
	Cassandra *gocql.Session
}
