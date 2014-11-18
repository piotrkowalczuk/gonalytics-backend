package v1

import (
	"github.com/gocraft/web"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
	"github.com/piotrkowalczuk/gonalytics-backend/lib/models"
	"github.com/piotrkowalczuk/gonalytics-backend/services"
	"strconv"
)

// Get ...
func (bc *BaseContext) VisitsGETHandler(w web.ResponseWriter, r *web.Request) {
	r.ParseForm()

	siteID, err := strconv.ParseInt(r.FormValue("t.sid"), 10, 64)
	if err != nil {
		bc.HTTPError(w, err, "Unexpected error.", http.StatusBadRequest)
		return
	}

	deviceIsTablet, _ := strconv.ParseBool(r.FormValue("d.it"))
	deviceIsPhone, _ := strconv.ParseBool(r.FormValue("d.ip"))
	deviceIsMobile, _ := strconv.ParseBool(r.FormValue("d.im"))
	browserPluginJava, _ := strconv.ParseBool(r.FormValue("b.p.j"))
	browserCookie, _ := strconv.ParseBool(r.FormValue("b.c"))
	browserIsOnline, _ := strconv.ParseBool(r.FormValue("b.io"))
	browserWindowWidth, _ := strconv.ParseInt(r.FormValue("b.w.w"), 10, 64)
	browserWindowHeight, _ := strconv.ParseInt(r.FormValue("b.w.h"), 10, 64)
	screenWidth, _ := strconv.ParseInt(r.FormValue("s.w"), 10, 64)
	screenHeight, _ := strconv.ParseInt(r.FormValue("s.h"), 10, 64)

	trackRequest := models.TrackRequest{
		SiteID:                 int64(siteID),
		RemoteAddress:          r.RemoteAddr,
		Domain:                 r.Header.Get("Origin"),
		VisitID:                r.FormValue("v.id"),
		PageTitle:              r.FormValue("p.t"),
		PageHost:               r.FormValue("p.h"),
		PageURL:                r.FormValue("p.u"),
		Language:               r.FormValue("lng"),
		Referrer:               r.FormValue("r"),
		BrowserPluginJava:      browserPluginJava,
		BrowserName:            r.FormValue("b.n"),
		BrowserVersion:         r.FormValue("b.v"),
		BrowserMajorVersion:    r.FormValue("b.mv"),
		BrowserUserAgent:       r.UserAgent(),
		BrowserPlatform:        r.FormValue("b.p"),
		BrowserCookie:          browserCookie,
		BrowserIsOnline:        browserIsOnline,
		BrowserWindowWidth:     int64(browserWindowWidth),
		BrowserWindowHeight:    int64(browserWindowHeight),
		OperatingSystemName:    r.FormValue("os.n"),
		OperatingSystemVersion: r.FormValue("os.v"),
		ScreenWidth:            int64(screenWidth),
		ScreenHeight:           int64(screenHeight),
		DeviceName:             r.FormValue("d.n"),
		DeviceIsTablet:         deviceIsTablet,
		DeviceIsPhone:          deviceIsPhone,
		DeviceIsMobile:         deviceIsMobile,
		MadeAt:                 time.Now(),
	}

	isNewVisit := trackRequest.IsNewVisit()

	actionCreator := lib.NewActionCreator(services.GeoIP)
	action, err := actionCreator.Create(&trackRequest)
	if err != nil {
		bc.HTTPError(w, err, "Unexpected error.", http.StatusInternalServerError)
		return
	}

	err = bc.RepositoryManager.Action.Insert(action)
	if err != nil {
		bc.HTTPError(w, err, "Unexpected error.", http.StatusInternalServerError)
		return
	}

	trackRequest.VisitID = action.VisitID.String()

	message := ""
	if isNewVisit {
		message = "New visit"
	} else {
		message = "Existing visit"
	}

	bc.Logger.WithFields(logrus.Fields{
		"url":     trackRequest.PageURL,
		"visitId": trackRequest.VisitID,
	}).Info(message)

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Gonalytics-Visit-Id")
	w.Header().Set("Access-Control-Allow-Origin", trackRequest.Domain)
	w.Header().Set("Gonalytics-Visit-Id", trackRequest.VisitID)
	http.ServeFile(w, r.Request, "data/1x1.gif")
}
