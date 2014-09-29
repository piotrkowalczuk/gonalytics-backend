package v1

import (
	"net/http"
	"time"

	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/services"
	"labix.org/v2/mgo/bson"
)

// VisitController ...
type VisitController struct {
	BaseController
}

// Get ...
func (vc *VisitController) Get() {
	w := vc.Ctx.ResponseWriter
	r := vc.Ctx.Request

	siteID, err := vc.GetInt("t.sid")
	vc.AbortIf(err, "Unexpected error.", http.StatusBadRequest)

	now := time.Now()
	mongoDateNow := models.NewMongoDate(&now)

	deviceIsTablet, _ := vc.GetBool("d.it")
	deviceIsPhone, _ := vc.GetBool("d.ip")
	deviceIsMobile, _ := vc.GetBool("d.im")
	browserPluginJava, _ := vc.GetBool("b.p.j")
	browserCookie, _ := vc.GetBool("b.c")
	browserIsOnline, _ := vc.GetBool("b.io")
	browserWindowWidth, _ := vc.GetInt("b.w.w")
	browserWindowHeight, _ := vc.GetInt("b.w.h")
	screenWidth, _ := vc.GetInt("s.w")
	screenHeight, _ := vc.GetInt("s.h")

	trackRequest := models.TrackRequest{
		SiteID:                 siteID,
		RemoteAddress:          r.RemoteAddr,
		Domain:                 r.Header.Get("Origin"),
		VisitID:                vc.GetString("v.id"),
		PageTitle:              vc.GetString("p.t"),
		PageHost:               vc.GetString("p.h"),
		PageURL:                vc.GetString("p.u"),
		Language:               vc.GetString("lng"),
		Referrer:               vc.GetString("r"),
		BrowserPluginJava:      browserPluginJava,
		BrowserName:            vc.GetString("b.n"),
		BrowserVersion:         vc.GetString("b.v"),
		BrowserMajorVersion:    vc.GetString("b.mv"),
		BrowserUserAgent:       r.UserAgent(),
		BrowserPlatform:        vc.GetString("b.p"),
		BrowserCookie:          browserCookie,
		BrowserIsOnline:        browserIsOnline,
		BrowserWindowWidth:     browserWindowWidth,
		BrowserWindowHeight:    browserWindowHeight,
		OperatingSystemName:    vc.GetString("os.n"),
		OperatingSystemVersion: vc.GetString("os.v"),
		ScreenWidth:            screenWidth,
		ScreenHeight:           screenHeight,
		DeviceName:             vc.GetString("d.n"),
		DeviceIsTablet:         deviceIsTablet,
		DeviceIsPhone:          deviceIsPhone,
		DeviceIsMobile:         deviceIsMobile,
		MadeAt:                 mongoDateNow.DateTime,
		MadeAtBucket:           mongoDateNow.Bucket,
	}

	if trackRequest.IsNewVisit() {
		vc.Log.Debug("New visit")

		visivcreator := services.NewVisitCreator(&trackRequest)
		err = vc.MongoDB.C("visit").Insert(&visivcreator.Visit)
		trackRequest.VisitID = visivcreator.Visit.ID.Hex()

		actionCreator := services.NewActionCreator(&trackRequest)
		err = vc.MongoDB.C("action").Insert(&actionCreator.Action)

		vc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	} else {
		vc.Log.Debug("Existing visit #%s", trackRequest.VisitID)

		actionCreator := services.NewActionCreator(&trackRequest)

		err = vc.MongoDB.C("action").Insert(&actionCreator.Action)
		err = vc.MongoDB.C("visit").UpdateId(
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

		vc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Gonalytics-Visit-Id")
	w.Header().Set("Access-Control-Allow-Origin", trackRequest.Domain)
	w.Header().Set("Gonalytics-Visit-Id", trackRequest.VisitID)
	http.ServeFile(w, r, "data/1x1.gif")
}
