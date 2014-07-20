package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"github.com/piotrkowalczuk/gowik-tracker/services"
	"labix.org/v2/mgo/bson"
	"net/http"
	"time"
)

type TrackController struct {
	BaseController
}

func (tc *TrackController) Get() {
	var err error

	w := tc.Ctx.ResponseWriter
	r := tc.Ctx.Request
	now := time.Now()
	visit := models.Visit{}
	domain := r.Header.Get("Origin")
	visitId := tc.GetString("v.id")

	if len(visitId) == 0 {
		tc.log.Debug("New visit")
		visit.Id = bson.NewObjectId()
		visit.CreatedAt = models.NewMongoDate(now)
		tc.MongoPool.Collection("visit").Insert(&visit)
	} else {
		tc.log.Debug("Existing visit #%s", visitId)
		tc.MongoPool.Collection("visit").FindId(bson.ObjectIdHex(visitId)).One(&visit)
	}

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

	website := models.Website{
		Title: tc.GetString("w.t"),
		Host:  tc.GetString("w.h"),
		Url:   domain,
	}

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

	action := models.Action{
		Id:              bson.NewObjectId(),
		VisitId:         visit.Id,
		Referrer:        tc.GetString("r"),
		Language:        tc.GetString("lng"),
		CreatedAt:       models.NewMongoDate(now),
		Browser:         &browser,
		Website:         &website,
		OperatingSystem: &os,
		Screen:          &screen,
		Device:          &device,
	}

	geoLocation, err := services.NewGeoLocation("80.48.120.255")
	if err == nil {
		location := models.NewLocationFromGeoIP(geoLocation.Location)
		action.Location = location
	}

	tc.MongoPool.Collection("action").Insert(&action)
	tc.MongoPool.Collection("visit").UpdateId(visit.Id, bson.M{"$push": bson.M{"actions": action.Id}})

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Expose-Headers", "Gowik-Visit-Id")
	w.Header().Set("Access-Control-Allow-Origin", domain)
	w.Header().Set("Gowik-Visit-Id", visit.Id.Hex())
	http.ServeFile(w, r, "1x1.gif")
}
