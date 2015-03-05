package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gocql/gocql"
	"github.com/gocraft/web"
	"github.com/piotrkowalczuk/gonalytics-backend/lib"
)

// VisitsGETHandler ...
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

	trackRequest := lib.TrackRequest{
		SiteID:                 int64(siteID),
		RemoteAddress:          r.RemoteAddr,
		Domain:                 r.Header.Get("Origin"),
		VisitID:                r.FormValue("v.id"),
		InitializeVisit:        false,
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
		ServerTime:             time.Now().UTC(),
	}

	message := ""
	incomingVisitID := trackRequest.VisitID

	if trackRequest.IsValidVisitID() {
		message = "Valid visit id"

		visitIDUUID, err := trackRequest.ParseVisitID()
		if err != nil {
			bc.HTTPError(w, err, "Unexpected error.", http.StatusInternalServerError)
			return
		}

		isActiveVisit, err := bc.RepositoryManager.
			VisitActions.
			IsActiveVisit(visitIDUUID)
		if err != nil {
			bc.HTTPError(w, err, "Unexpected error.", http.StatusInternalServerError)
			return
		}

		if !isActiveVisit {
			trackRequest.InitializeVisit = true
			trackRequest.VisitID = gocql.TimeUUID().String()
			message = "Outdated visit id"
		}
	} else {
		trackRequest.VisitID = gocql.TimeUUID().String()
		trackRequest.InitializeVisit = true
		message = "Missing or invalid visit id"
	}

	trackRequestBytes, err := json.Marshal(trackRequest)
	if err != nil {
		bc.HTTPError(w, err, "Unexpected error.", http.StatusInternalServerError)
		return
	}
	err = bc.KafkaPublisher.PublishAction(string(trackRequestBytes))
	if err != nil {
		bc.HTTPError(w, err, "Unexpected error.", http.StatusInternalServerError)
		return
	}

	bc.Logger.WithFields(logrus.Fields{
		"url":             trackRequest.PageURL,
		"incomingVisitID": incomingVisitID,
		"outgoingVisitID": trackRequest.VisitID,
	}).Info(message)

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Gonalytics-Visit-Id")
	w.Header().Set("Access-Control-Allow-Origin", trackRequest.Domain)
	w.Header().Set("Gonalytics-Visit-Id", trackRequest.VisitID)
	http.ServeFile(w, r.Request, "data/1x1.gif")
}
