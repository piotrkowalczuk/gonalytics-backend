package reports

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"labix.org/v2/mgo/cql"
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
			Id:              cql.NewObjectId(),
			Referrer:        "r",
			Page:            &page,
			CreatedAt:       mongoDateNow.DateTime,
			CreatedAtBucket: mongoDateNow.Bucket,
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
			Id:                  cql.NewObjectId(),
			Referrer:            "r",
			Language:            "lng",
			Actions:             []*Action{&action},
			NbOfActions:         1,
			Location:            &location,
			Browser:             &browser,
			FirstPage:           &page,
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
		visits[i] = &visit
	}

	return visits
}

func benchmarkNewCountryreportsFromVisits(b *testing.B, nbOfVisits int) {
	visits := generateVisists(nbOfVisits)
	b.StopTimer()
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewCountryreportsFromVisits(visits)
	}
}

func BenchmarkNewCountryreportsFromVisitsSmall(b *testing.B) {
	benchmarkNewCountryreportsFromVisits(b, 100)
}

func BenchmarkNewCountryreportsFromVisitsMedium(b *testing.B) {
	benchmarkNewCountryreportsFromVisits(b, 1000)
}

func BenchmarkNewCountryreportsFromVisitsLarge(b *testing.B) {
	benchmarkNewCountryreportsFromVisits(b, 10000)
}

func BenchmarkNewCountryreportsFromVisitsHuge(b *testing.B) {
	benchmarkNewCountryreportsFromVisits(b, 100000)
}
