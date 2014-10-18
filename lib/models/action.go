package models

import (
	"time"

	"github.com/gocql/gocql"
)

// Action represents single action of visitor.
type Action struct {
	ID        gocql.UUID `json:"id" cql:"id,omitempty"`
	VisitID   gocql.UUID `json:"visitId" cql:"visit_id"`
	Referrer  string     `json:"referrer" cql:"referrer"`
	Page      *Page      `json:"page" cql:"page"`
	CreatedAt time.Time  `json:"createdAt" cql:"created_at"`
}

// Actions is a simple wrapper for slice of actions
type Actions []*Action

// Append add new action at the end of collection
func (a *Actions) Append(action *Action) {
	*a = append(*a, action)
}
