package services

import (
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"labix.org/v2/mgo/bson"
)

// ActionCreator ...
type ActionCreator struct {
	trackRequest *models.TrackRequest
	Action       *models.Action
}

// NewActionCreator ...
func NewActionCreator(trackRequest *models.TrackRequest) *ActionCreator {
	ac := ActionCreator{
		trackRequest: trackRequest,
	}

	ac.createAction()

	return &ac
}

func (ac *ActionCreator) createAction() {
	ac.Action = &models.Action{
		ID:        bson.NewObjectId(),
		VisitID:   bson.ObjectIdHex(ac.trackRequest.VisitID),
		Referrer:  ac.trackRequest.Referrer,
		Page:      ac.createPage(),
		CreatedAt: ac.trackRequest.MadeAt,
	}
}

func (ac *ActionCreator) createPage() *models.Page {
	return &models.Page{
		Title: ac.trackRequest.PageTitle,
		Host:  ac.trackRequest.PageHost,
		Url:   ac.trackRequest.PageURL,
	}
}
