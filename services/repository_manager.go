package services

import (
	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"
)

// Singleton instance of RepositoryManager.
var RepositoryManager lib.RepositoryManager

// InitRepositoryManager ...
func InitRepositoryManager(cassandra *gocql.Session) {
	repository := repositories.Repository{
		Cassandra: cassandra,
	}

	RepositoryManager = lib.RepositoryManager{
		Visit: repositories.VisitRepository{
			Repository: repository,
		},
	}
}
