package lib

import "github.com/piotrkowalczuk/gonalytics-backend/lib/repositories"

// RepositoryManager ...
type RepositoryManager struct {
	Visit  repositories.VisitRepository
	Action repositories.ActionRepository
}
