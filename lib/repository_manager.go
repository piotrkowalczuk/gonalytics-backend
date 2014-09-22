package lib

import "github.com/piotrkowalczuk/gonalytics-tracker/repositories"

// RepositoryManager ...
type RepositoryManager struct {
	Visit  repositories.VisitRepository
	Action repositories.ActionRepository
}
