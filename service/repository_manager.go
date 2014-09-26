package service

import (
	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"
	"labix.org/v2/mgo"
)

// RepositoryManager ...
var RepositoryManager lib.RepositoryManager

// InitRepositoryManager ...
func InitRepositoryManager(MongoDB *mgo.Database, Cassandra *gocql.Session) {
	repository := repositories.Repository{
		MongoDB:   MongoDB,
		Cassandra: Cassandra,
	}

	RepositoryManager = lib.RepositoryManager{
		Visit:  repositories.VisitRepository{repository},
		Action: repositories.ActionRepository{repository},
	}
}
