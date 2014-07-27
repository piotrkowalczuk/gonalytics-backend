package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"github.com/piotrkowalczuk/gowik-tracker/services"
	"labix.org/v2/mgo/bson"
	"net/http"
	"log"
	"net"
	"time"
)

type TrackController struct {
	BaseController
}

func (tc *TrackController) Get() {
	var err error

	siteId, err := tc.GetInt("t.sid")
	tc.abortIf(err, http.StatusBadRequest)

	w := tc.Ctx.ResponseWriter
	r := tc.Ctx.Request
	mongoDateNow := models.NewMongoDate(time.Now())
	domain := r.Header.Get("Origin")
	visitId := tc.GetString("v.id")
	requestIp, _, _ := net.SplitHostPort(r.RemoteAddr)
	var visit models.Visit

	page := models.Page{
		Title: tc.GetString("p.t"),
		Host:  tc.GetString("p.h"),
		Url:   tc.GetString("p.u"),
	}

	action := models.Action{
		Id:              bson.NewObjectId(),
		Referrer:        tc.GetString("r"),
		Page:            &page,
		CreatedAt:       mongoDateNow.DateTime,
		CreatedAtBucket: mongoDateNow.Bucket,
	}

	if len(visitId) == 0 {
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

		geoLocation, err := services.NewGeoLocation(requestIp)
		location := models.Location{}
		if err == nil {
			location = *models.NewLocationFromGeoIP(geoLocation.Location)
		}

		visit = models.Visit{
			Id:                  bson.NewObjectId(),
			Referrer:            tc.GetString("r"),
			Language:            tc.GetString("lng"),
			Actions:             []*models.Action{&action},
			NbOfActions:         1,
			SiteId:				 siteId,
			Location:            &location,
			Browser:             &browser,
			FirstPage:           &page,
			LastPage:            &page,
			OperatingSystem:     &os,
			Screen:              &screen,
			Device:              &device,
			CreatedAt:           mongoDateNow.DateTime,
			CreatedAtBucket:     mongoDateNow.Bucket,
			FirstActionAt:       mongoDateNow.DateTime,
			FirstActionAtBucket: mongoDateNow.Bucket,
			LastActionAt:        mongoDateNow.DateTime,
			LastActionAtBucket:  mongoDateNow.Bucket,
		}

		visitId = visit.Id.Hex()
		err = tc.MongoPool.Collection("visit").Insert(&visit)
		tc.abortIf(err, http.StatusInternalServerError)
	} else {
		tc.log.Debug("Existing visit #%s", visitId)

		err = tc.MongoPool.Collection("visit").UpdateId(
			bson.ObjectIdHex(visitId),
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

		tc.abortIf(err, http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Gowik-Visit-Id")
	w.Header().Set("Access-Control-Allow-Origin", domain)
	w.Header().Set("Gowik-Visit-Id", visitId)
	http.ServeFile(w, r, "1x1.gif")
}
