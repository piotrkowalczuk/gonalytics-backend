package lib

import (
	"net"
	"strconv"
	"time"

	"github.com/gocql/gocql"
)

// TrackRequest ...
type TrackRequest struct {
	SiteID                 int64
	RemoteAddress          string
	Domain                 string
	VisitID                string
	InitializeVisit        bool
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

// CreateMap ...
func (tr *TrackRequest) CreateMap() map[string]string {
	return map[string]string{
		"siteId":                 strconv.FormatInt(tr.SiteID, 10),
		"visitId":                tr.VisitID,
		"remoteAddress":          tr.RemoteAddress,
		"domain":                 tr.Domain,
		"pageTitle":              tr.PageTitle,
		"pageHost":               tr.PageHost,
		"pageURL":                tr.PageURL,
		"language":               tr.Language,
		"referrer":               tr.Referrer,
		"browserPluginJava":      strconv.FormatBool(tr.BrowserPluginJava),
		"browserName":            tr.BrowserName,
		"browserVersion":         tr.BrowserVersion,
		"browserMajorVersion":    tr.BrowserMajorVersion,
		"browserUserAgent":       tr.BrowserUserAgent,
		"browserPlatform":        tr.BrowserPlatform,
		"browserCookie":          strconv.FormatBool(tr.BrowserCookie),
		"browserIsOnline":        strconv.FormatBool(tr.BrowserIsOnline),
		"browserWindowWidth":     strconv.FormatInt(tr.BrowserWindowWidth, 10),
		"browserWindowHeight":    strconv.FormatInt(tr.BrowserWindowHeight, 10),
		"operatingSystemName":    tr.OperatingSystemName,
		"operatingSystemVersion": tr.OperatingSystemVersion,
		"screenWidth":            strconv.FormatInt(tr.ScreenWidth, 10),
		"screenHeight":           strconv.FormatInt(tr.ScreenHeight, 10),
		"deviceName":             tr.DeviceName,
		"deviceIsTablet":         strconv.FormatBool(tr.DeviceIsTablet),
		"deviceIsPhone":          strconv.FormatBool(tr.DeviceIsPhone),
		"deviceIsMobile":         strconv.FormatBool(tr.DeviceIsMobile),
	}
}

// GetRequestIP ...
func (tr *TrackRequest) GetRequestIP() string {
	requestIP, _, _ := net.SplitHostPort(tr.RemoteAddress)

	return requestIP
}

// IsValidVisitID ...
func (tr *TrackRequest) IsValidVisitID() bool {
	if len(tr.VisitID) == 0 {
		return false
	}

	if _, err := gocql.ParseUUID(tr.VisitID); err != nil {
		return false
	}

	return true
}

// ParseVisitID ...
func (tr *TrackRequest) ParseVisitID() (gocql.UUID, error) {
	return gocql.ParseUUID(tr.VisitID)
}
