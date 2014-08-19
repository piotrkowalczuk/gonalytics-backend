package controllers

import (
	"net/http"
	"time"

	"github.com/piotrkowalczuk/gonalytics-tracker/models"
	"github.com/piotrkowalczuk/gonalytics-tracker/services"
	"github.com/piotrkowalczuk/gonalytics-tracker/structs"
	"labix.org/v2/mgo/bson"
)

// TrackController ...
type TrackController struct {
	BaseController
}

// Get ...
func (tc *TrackController) Get() {
	var err error

	w := tc.Ctx.ResponseWriter
	r := tc.Ctx.Request

	siteID, err := tc.GetInt("t.sid")
	tc.AbortIf(err, "Unexpected error.", http.StatusBadRequest)

	now := time.Now()
	mongoDateNow := models.NewMongoDate(&now)

	deviceIsTablet, _ := tc.GetBool("d.it")
	deviceIsPhone, _ := tc.GetBool("d.ip")
	deviceIsMobile, _ := tc.GetBool("d.im")
	browserPluginJava, _ := tc.GetBool("b.p.j")
	browserCookie, _ := tc.GetBool("b.c")
	browserIsOnline, _ := tc.GetBool("b.io")
	browserWindowWidth, _ := tc.GetInt("b.w.w")
	browserWindowHeight, _ := tc.GetInt("b.w.h")
	screenWidth, _ := tc.GetInt("s.w")
	screenHeight, _ := tc.GetInt("s.h")

	trackRequest := structs.TrackRequest{
		SiteID:                 siteID,
		RemoteAddress:          r.RemoteAddr,
		Domain:                 r.Header.Get("Origin"),
		VisitID:                tc.GetString("v.id"),
		PageTitle:              tc.GetString("p.t"),
		PageHost:               tc.GetString("p.h"),
		PageURL:                tc.GetString("p.u"),
		Language:               tc.GetString("lng"),
		Referrer:               tc.GetString("r"),
		BrowserPluginJava:      browserPluginJava,
		BrowserName:            tc.GetString("b.n"),
		BrowserVersion:         tc.GetString("b.v"),
		BrowserMajorVersion:    tc.GetString("b.mv"),
		BrowserUserAgent:       r.UserAgent(),
		BrowserPlatform:        tc.GetString("b.p"),
		BrowserCookie:          browserCookie,
		BrowserIsOnline:        browserIsOnline,
		BrowserWindowWidth:     browserWindowWidth,
		BrowserWindowHeight:    browserWindowHeight,
		OperatingSystemName:    tc.GetString("os.n"),
		OperatingSystemVersion: tc.GetString("os.v"),
		ScreenWidth:            screenWidth,
		ScreenHeight:           screenHeight,
		DeviceName:             tc.GetString("d.n"),
		DeviceIsTablet:         deviceIsTablet,
		DeviceIsPhone:          deviceIsPhone,
		DeviceIsMobile:         deviceIsMobile,
		MadeAt:                 mongoDateNow.DateTime,
		MadeAtBucket:           mongoDateNow.Bucket,
	}

	var visitID string
	visitID = trackRequest.VisitID

	if trackRequest.IsNewVisit() {
		tc.log.Debug("New visit")

		visitCreator := services.NewVisitCreator(&trackRequest)

		err = tc.MongoPool.Collection("visit").Insert(&visitCreator.Visit)
		visitID = visitCreator.Visit.ID.Hex()
		trackRequest.VisitID = visitID

		actionCreator := services.NewActionCreator(&trackRequest)
		err = tc.MongoPool.Collection("action").Insert(&actionCreator.Action)

		tc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	} else {
		tc.log.Debug("Existing visit #%s", visitID)

		actionCreator := services.NewActionCreator(&trackRequest)

		err = tc.MongoPool.Collection("action").Insert(&actionCreator.Action)
		err = tc.MongoPool.Collection("visit").UpdateId(
			bson.ObjectIdHex(visitID),
			bson.M{
				"$inc": bson.M{"nb_of_actions": 1},
				"$set": bson.M{
					"last_action_at":        trackRequest.MadeAt,
					"last_action_at_bucket": trackRequest.MadeAtBucket,
					"last_page":             &actionCreator.Action.Page,
				},
			},
		)

		tc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Gonalytics-Visit-Id")
	w.Header().Set("Access-Control-Allow-Origin", trackRequest.Domain)
	w.Header().Set("Gonalytics-Visit-Id", visitID)
	http.ServeFile(w, r, "data/1x1.gif")
}
