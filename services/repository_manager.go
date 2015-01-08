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
		VisitAction: repositories.VisitActionRepository{
			Repository: repository,
		},
		SiteDayCountryActionsCounter: repositories.SiteDayCountryActionsCounterRepository{
			Repository: repository,
		},
		SiteMonthCountryActionsCounter: repositories.SiteMonthCountryActionsCounterRepository{
			Repository: repository,
		},
		SiteYearCountryActionsCounter: repositories.SiteYearCountryActionsCounterRepository{
			Repository: repository,
		},
	}
}
