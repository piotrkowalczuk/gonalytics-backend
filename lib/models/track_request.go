package models

import (
	"net"
	"time"
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
	MadeAt                 *time.Time
	MadeAtBucket           []string
}

// GetRequestIP ...
func (tr *TrackRequest) GetRequestIP() string {
	requestIP, _, _ := net.SplitHostPort(tr.RemoteAddress)

	return requestIP
}

// IsNewVisit ...
func (tr *TrackRequest) IsNewVisit() bool {
	return len(tr.VisitID) == 0
}
