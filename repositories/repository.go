package repositories

import (
	"labix.org/v2/mgo"
)

// Repository ...
type Repository struct {
	MongoDB *mgo.Database
}
