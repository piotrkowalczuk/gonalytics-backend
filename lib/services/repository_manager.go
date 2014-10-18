package services

import (
	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"
)

// RepositoryManager ...
var RepositoryManager lib.RepositoryManager

// InitRepositoryManager ...
func InitRepositoryManager(Cassandra *gocql.Session) {
	repository := repositories.Repository{
		Cassandra: Cassandra,
	}

	RepositoryManager = lib.RepositoryManager{
		Visit: repositories.VisitRepository{repository},
	}
}
