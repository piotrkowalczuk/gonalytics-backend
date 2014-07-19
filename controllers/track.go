package controllers

import (
	"github.com/piotrkowalczuk/gowik-tracker/models"
	"net/http"
	"time"
)

type TrackController struct {
	BaseController
}

func (tc *TrackController) Get() {
	w := tc.Ctx.ResponseWriter
	r := tc.Ctx.Request
	var userIdCookie *http.Cookie
	var err error

	userIdCookie, err = tc.Ctx.Request.Cookie("userId")

	if err != nil {
		userIdCookie = &http.Cookie{
			Name:   "userId",
			Value:  "test",
			Domain: "",
		}
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
		UserAgent:    tc.GetString("b.ua"),
		Platform:     tc.GetString("b.p"),
		Plugins:      plugins,
		Window:       window,
	}
	browser.Cookie, _ = tc.GetBool("b.c")
	browser.IsOnline, _ = tc.GetBool("b.io")

	website := models.Website{
		Title: tc.GetString("w.t"),
		Host:  tc.GetString("w.h"),
		Url:   tc.GetString("w.u"),
	}

	os := models.OperatingSystem{
		Name: tc.GetString("os.n"),
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
		UserId:          userIdCookie.Value,
		Referrer:        tc.GetString("r"),
		Language:        tc.GetString("lng"),
		CreatedAt:       models.NewMongoDate(time.Now()),
		Browser:         &browser,
		Website:         &website,
		OperatingSystem: &os,
		Screen:          &screen,
		Device:          &device,
	}

	tc.MongoPool.Collection("action").Insert(action)

	http.SetCookie(w, userIdCookie)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	http.ServeFile(w, r, "1x1.gif")
}
