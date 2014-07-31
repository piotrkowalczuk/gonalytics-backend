package controllers

import (
	"net"
	"net/http"
	"time"

	"github.com/piotrkowalczuk/gonalytics-tracker/models"
	"github.com/piotrkowalczuk/gonalytics-tracker/services"
	"labix.org/v2/mgo/bson"
)

// TrackController ...
type TrackController struct {
	BaseController
}

// Get ...
func (tc *TrackController) Get() {
	var err error

	siteID, err := tc.GetInt("t.sid")
	tc.AbortIf(err, "Unexpected error.", http.StatusBadRequest)

	w := tc.Ctx.ResponseWriter
	r := tc.Ctx.Request
	now := time.Now()
	mongoDateNow := models.NewMongoDate(&now)
	domain := r.Header.Get("Origin")
	visitID := tc.GetString("v.id")
	requestIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	var visit models.Visit

	page := models.Page{
		Title: tc.GetString("p.t"),
		Host:  tc.GetString("p.h"),
		Url:   tc.GetString("p.u"),
	}

	action := models.Action{
		ID:              bson.NewObjectId(),
		Referrer:        tc.GetString("r"),
		Page:            &page,
		CreatedAt:       mongoDateNow.DateTime,
		CreatedAtBucket: mongoDateNow.Bucket,
	}

	if len(visitID) == 0 {
		tc.log.Debug("New visit")

		plugins := models.Plugins{}
		plugins.Java, _ = tc.GetBool("b.p.j")

		window := models.Window{}
		window.Width, _ = tc.GetInt("b.w.w")
		window.Height, _ = tc.GetInt("b.w.h")

		browser := models.Browser{
			Name:         tc.GetString("b.n"),
			Version:      tc.GetString("b.v"),
			MajorVersion: tc.GetString("b.mv"),
			UserAgent:    r.UserAgent(),
			Platform:     tc.GetString("b.p"),
			Plugins:      plugins,
			Window:       window,
		}
		browser.Cookie, _ = tc.GetBool("b.c")
		browser.IsOnline, _ = tc.GetBool("b.io")

		os := models.OperatingSystem{
			Name:    tc.GetString("os.n"),
			Version: tc.GetString("os.v"),
		}

		screen := models.Screen{}
		screen.Width, _ = tc.GetInt("s.w")
		screen.Height, _ = tc.GetInt("s.h")

		device := models.Device{
			Name: tc.GetString("d.n"),
		}
		device.IsTablet, _ = tc.GetBool("d.it")
		device.IsPhone, _ = tc.GetBool("d.ip")
		device.IsMobile, _ = tc.GetBool("d.im")

		geoLocation, err := services.NewGeoLocation(requestIP)
		location := models.Location{}
		if err == nil {
			location = *models.NewLocationFromGeoIP(geoLocation.Location)
		}

		visit = models.Visit{
			ID:                  bson.NewObjectId(),
			Referrer:            tc.GetString("r"),
			Language:            tc.GetString("lng"),
			Actions:             []*models.Action{&action},
			NbOfActions:         1,
			SiteID:              siteID,
			Location:            &location,
			Browser:             &browser,
			FirstPage:           &page,
			LastPage:            &page,
			OperatingSystem:     &os,
			Screen:              &screen,
			Device:              &device,
			FirstActionAt:       mongoDateNow.DateTime,
			FirstActionAtBucket: mongoDateNow.Bucket,
			LastActionAt:        mongoDateNow.DateTime,
			LastActionAtBucket:  mongoDateNow.Bucket,
		}

		visitID = visit.ID.Hex()
		err = tc.MongoPool.Collection("visit").Insert(&visit)
		tc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	} else {
		tc.log.Debug("Existing visit #%s", visitID)

		err = tc.MongoPool.Collection("visit").UpdateId(
			bson.ObjectIdHex(visitID),
			bson.M{
				"$push": bson.M{"actions": action},
				"$inc":  bson.M{"nb_of_actions": 1},
				"$set": bson.M{
					"last_action_at":        mongoDateNow.DateTime,
					"last_action_at_bucket": mongoDateNow.Bucket,
					"last_page":             &page,
				},
			},
		)

		tc.AbortIf(err, "Unexpected error.", http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Gowik-Visit-Id")
	w.Header().Set("Access-Control-Allow-Origin", domain)
	w.Header().Set("Gowik-Visit-Id", visitID)
	http.ServeFile(w, r, "1x1.gif")
}
