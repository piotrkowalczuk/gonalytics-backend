package services

import (
	"github.com/gocql/gocql"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
)

// ActionCreator ...
type ActionCreator struct {
	trackRequest *models.TrackRequest
	Action       *models.Action
}

// NewActionCreator ...
func NewActionCreator(trackRequest *models.TrackRequest) (*ActionCreator, error) {
	ac := ActionCreator{
		trackRequest: trackRequest,
	}

	if err := ac.createAction(); err != nil {
		return nil, err
	}

	return &ac, nil
}

func (ac *ActionCreator) createAction() error {
	visitID, err := gocql.ParseUUID(ac.trackRequest.VisitID)

	if err != nil {
		return err
	}

	ac.Action = &models.Action{
		ID:        gocql.TimeUUID(),
		VisitID:   visitID,
		Referrer:  ac.trackRequest.Referrer,
		Page:      ac.createPage(),
		CreatedAt: ac.trackRequest.MadeAt,
	}

	return nil
}

func (ac *ActionCreator) createPage() *models.Page {
	return &models.Page{
		Title: ac.trackRequest.PageTitle,
		Host:  ac.trackRequest.PageHost,
		URL:   ac.trackRequest.PageURL,
	}
}
