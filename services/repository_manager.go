package services

import (
	"github.com/piotrkowalczuk/gonalytics-tracker/lib"
	"github.com/piotrkowalczuk/gonalytics-tracker/repositories"
	"labix.org/v2/mgo"
)

// RepositoryManager ...
var RepositoryManager lib.RepositoryManager

// InitRepositoryManager ...
func InitRepositoryManager(MongoDB *mgo.Database) {
	repository := repositories.Repository{
		MongoDB: MongoDB,
	}

	RepositoryManager = lib.RepositoryManager{
		Visit:  repositories.VisitRepository{repository},
		Action: repositories.ActionRepository{repository},
	}
}
