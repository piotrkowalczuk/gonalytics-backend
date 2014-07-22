package models

import (
	"labix.org/v2/mgo/bson"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func generateVisists(nbOfVisits int) []*Visit {
	visits := make([]*Visit, nbOfVisits)

	for i := 0; i < nbOfVisits; i++ {
		mongoDateNow := NewMongoDate(time.Now())

		page := Page{
			Title: "p.t",
			Host:  "p.h",
			Url:   "p.u",
		}

		action := Action{
			Id:        bson.NewObjectId(),
			Referrer:  "r",
			Page:      &page,
			CreatedAt: mongoDateNow,
		}

		plugins := Plugins{}
		plugins.Java = true

		window := Window{}
		window.Width = 1000
		window.Height = 1000

		browser := Browser{
			Name:         "b.n",
			Version:      "b.v",
			MajorVersion: "b.mv",
			UserAgent:    "user-agent",
			Platform:     "b.p",
			Plugins:      plugins,
			Window:       window,
		}
		browser.Cookie = true
		browser.IsOnline = true

		os := OperatingSystem{
			Name:    "os.n",
			Version: "os.v",
		}

		screen := Screen{}
		screen.Width = 1000
		screen.Height = 1000

		device := Device{
			Name: "d.n",
		}
		device.IsTablet = false
		device.IsPhone = false
		device.IsMobile = false

		rndId := rand.Intn(253)

		location := Location{
			CountryName: strconv.FormatInt(int64(rndId), 10),
			CountryCode: strconv.FormatInt(int64(rndId), 10),
			CountryId:   uint(rndId),
		}

		visit := Visit{
			Id:              bson.NewObjectId(),
			Referrer:        "r",
			Language:        "lng",
			Actions:         []*Action{&action},
			NbOfActions:     1,
			Location:        &location,
			Browser:         &browser,
			FirstPage:       &page,
			OperatingSystem: &os,
			Screen:          &screen,
			Device:          &device,
			CreatedAt:       mongoDateNow,
			FirstActionAt:   mongoDateNow,
			LastActionAt:    mongoDateNow,
		}
		visits[i] = &visit
	}

	return visits
}

func benchmarkNewCountryReportFromVisits(b *testing.B, nbOfVisits int) {
	visits := generateVisists(nbOfVisits)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewCountryReportFromVisits(visits)
	}
}

func BenchmarkNewCountryReportFromVisitsSmall(b *testing.B) {
	benchmarkNewCountryReportFromVisits(b, 100)
}

func BenchmarkNewCountryReportFromVisitsMedium(b *testing.B) {
	benchmarkNewCountryReportFromVisits(b, 1000)
}

func BenchmarkNewCountryReportFromVisitsLarge(b *testing.B) {
	benchmarkNewCountryReportFromVisits(b, 10000)
}

func BenchmarkNewCountryReportFromVisitsHuge(b *testing.B) {
	benchmarkNewCountryReportFromVisits(b, 100000)
}
