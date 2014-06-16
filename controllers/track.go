package controllers

import (
	"github.com/piotrkowalczuk/gonalytics/models"
	"net/http"
)

type TrackController struct {
	BaseController
}

func (tc *TrackController) Get() {
	report := new(models.Report)
	report.Name = tc.GetString("n")
	report.AppName = tc.GetString("an")
	report.Referrer = tc.GetString("r")
	report.Language = tc.GetString("lng")
	report.Cookie = tc.GetString("c")
	report.UserAgent = tc.GetString("ua")
	report.Java, _ = tc.GetBool("p.java")
	report.BrowserVersion = tc.GetString("b.v")
	report.BrowserVersionMinor = tc.GetString("b.vm")
	report.ScreenWidth, _ = tc.GetInt("s.w")
	report.ScreenHeight, _ = tc.GetInt("s.h")
	report.WebsiteTitle = tc.GetString("w.t")
	report.WebsiteHost = tc.GetString("w.h")
	report.WebsiteUrl = tc.GetString("w.u")

	tc.MongoPool.Collection("report").Insert(report)
	w := tc.Ctx.ResponseWriter
	r := tc.Ctx.Request
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, "1x1.gif")
}
