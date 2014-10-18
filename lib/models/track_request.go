package models

import (
	"net"
	"time"

	"github.com/gocql/gocql"
)

// TrackRequest ...
type TrackRequest struct {
	SiteID                 int64
	RemoteAddress          string
	Domain                 string
	VisitID                string
	PageTitle              string
	PageHost               string
	PageURL                string
	Language               string
	Referrer               string
	BrowserPluginJava      bool
	BrowserName            string
	BrowserVersion         string
	BrowserMajorVersion    string
	BrowserUserAgent       string
	BrowserPlatform        string
	BrowserCookie          bool
	BrowserIsOnline        bool
	BrowserWindowWidth     int64
	BrowserWindowHeight    int64
	OperatingSystemName    string
	OperatingSystemVersion string
	ScreenWidth            int64
	ScreenHeight           int64
	DeviceName             string
	DeviceIsTablet         bool
	DeviceIsPhone          bool
	DeviceIsMobile         bool
	MadeAt                 time.Time
}

// GetRequestIP ...
func (tr *TrackRequest) GetRequestIP() string {
	requestIP, _, _ := net.SplitHostPort(tr.RemoteAddress)

	return requestIP
}

// IsNewVisit ...
func (tr *TrackRequest) IsNewVisit() bool {
	if len(tr.VisitID) == 0 {
		return true
	}

	if _, err := gocql.ParseUUID(tr.VisitID); err != nil {
		return true
	}

	return false
}
