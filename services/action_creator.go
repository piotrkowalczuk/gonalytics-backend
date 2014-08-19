package services

import (
	"github.com/piotrkowalczuk/gonalytics-tracker/models"
	"github.com/piotrkowalczuk/gonalytics-tracker/structs"
	"labix.org/v2/mgo/bson"
)

// ActionCreator ...
type ActionCreator struct {
	trackRequest *structs.TrackRequest
	Action       *models.Action
}

// NewActionCreator ...
func NewActionCreator(trackRequest *structs.TrackRequest) *ActionCreator {
	ac := ActionCreator{
		trackRequest: trackRequest,
	}

	ac.createAction()

	return &ac
}

func (ac *ActionCreator) createAction() {
	ac.Action = &models.Action{
		ID:              bson.NewObjectId(),
		VisitID:         bson.ObjectIdHex(ac.trackRequest.VisitID),
		Referrer:        ac.trackRequest.Referrer,
		Page:            ac.createPage(),
		CreatedAt:       ac.trackRequest.MadeAt,
		CreatedAtBucket: ac.trackRequest.MadeAtBucket,
	}
}

func (ac *ActionCreator) createPage() *models.Page {
	return &models.Page{
		Title: ac.trackRequest.PageTitle,
		Host:  ac.trackRequest.PageHost,
		Url:   ac.trackRequest.PageURL,
	}
}
