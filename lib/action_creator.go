package lib

import (
	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

// ActionCreator ...
type ActionCreator struct {
	trackRequest *models.TrackRequest
}

// NewActionCreator ...
func NewActionCreator() *ActionCreator {
	return &ActionCreator{}
}

// Create ...
func (ac *ActionCreator) Create(trackRequest *models.TrackRequest) (*models.Action, error) {
	ac.trackRequest = trackRequest
	visitID, err := gocql.ParseUUID(ac.trackRequest.VisitID)

	if err != nil {
		return nil, err
	}

	return &models.Action{
		ID:        gocql.TimeUUID(),
		VisitID:   visitID,
		Referrer:  ac.trackRequest.Referrer,
		Page:      ac.createPage(),
		CreatedAt: ac.trackRequest.MadeAt,
	}, nil
}

func (ac *ActionCreator) createPage() *models.Page {
	return &models.Page{
		Title: ac.trackRequest.PageTitle,
		Host:  ac.trackRequest.PageHost,
		URL:   ac.trackRequest.PageURL,
	}
}
