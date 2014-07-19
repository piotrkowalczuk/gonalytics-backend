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

	action := new(models.Action)
	action.Name = tc.GetString("n")
	action.AppName = tc.GetString("an")
	action.Referrer = tc.GetString("r")
	action.Language = tc.GetString("lng")
	action.Cookie = tc.GetString("c")
	action.UserAgent = tc.GetString("ua")
	action.Java, _ = tc.GetBool("p.java")
	action.BrowserVersion = tc.GetString("b.v")
	action.BrowserVersionMinor = tc.GetString("b.vm")
	action.ScreenWidth, _ = tc.GetInt("s.w")
	action.ScreenHeight, _ = tc.GetInt("s.h")
	action.WebsiteTitle = tc.GetString("w.t")
	action.WebsiteHost = tc.GetString("w.h")
	action.WebsiteUrl = tc.GetString("w.u")
	action.UserId = userIdCookie.Value
	action.CreatedAt = models.NewMongoDate(time.Now())

	tc.MongoPool.Collection("report").Insert(action)

	http.SetCookie(w, userIdCookie)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	http.ServeFile(w, r, "1x1.gif")
}
