package v1

import (
	"net/http"
	"time"

	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"github.com/piotrkowalczuk/gonalytics-backend/service"
	"labix.org/v2/mgo/bson"
)

// TrackController ...
type TrackController struct {
	BaseController
}

// Get ...
func (tc *TrackController) Get() {
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

	trackRequest := models.TrackRequest{
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

	if trackRequest.IsNewVisit() {
		tc.Log.Debug("New visit")

		visitCreator := service.NewVisitCreator(&trackRequest)
		err = tc.MongoDB.C("visit").Insert(&visitCreator.Visit)
		trackRequest.VisitID = visitCreator.Visit.ID.Hex()

		actionCreator := service.NewActionCreator(&trackRequest)
		err = tc.MongoDB.C("action").Insert(&actionCreator.Action)

		tc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	} else {
		tc.Log.Debug("Existing visit #%s", trackRequest.VisitID)

		actionCreator := service.NewActionCreator(&trackRequest)

		err = tc.MongoDB.C("action").Insert(&actionCreator.Action)
		err = tc.MongoDB.C("visit").UpdateId(
			bson.ObjectIdHex(trackRequest.VisitID),
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
	w.Header().Set("Gonalytics-Visit-Id", trackRequest.VisitID)
	http.ServeFile(w, r, "data/1x1.gif")
}
