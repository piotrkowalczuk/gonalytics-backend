package models

import (
	"time"

	"github.com/gocql/gocql"
)

// ActionEntity represents visit object stored in database.
type ActionEntity struct {
	ID                          gocql.UUID    `cql:"id"`
	VisitID                     gocql.UUID    `cql:"visit_id"`
	IP                          string        `cql:"ip"`
	SiteID                      int64         `cql:"site_id"`
	Referrer                    string        `cql:"referrer"`
	Language                    string        `cql:"language"`
	BrowserName                 string        `cql:"browser_name"`
	BrowserVersion              string        `cql:"browser_version"`
	BrowserMajorVersion         string        `cql:"browser_major_version"`
	BrowserUserAgent            string        `cql:"browser_user_agent"`
	BrowserPlatform             string        `cql:"browser_platform"`
	BrowserCookie               bool          `cql:"browser_cookie"`
	BrowserPluginJava           bool          `cql:"browser_plugin_java"`
	BrowserIsOnline             bool          `cql:"browser_is_online"`
	BrowserWindowWidth          int64         `cql:"browser_window_width"`
	BrowserWindowHeight         int64         `cql:"browser_window_height"`
	ScreenWidth                 int64         `cql:"screen_width"`
	ScreenHeight                int64         `cql:"screen_height"`
	OperatingSystemName         string        `cql:"os_name"`
	OperatingSystemVersion      string        `cql:"os_version"`
	DeviceName                  string        `cql:"device_name"`
	DeviceIsMobile              bool          `cql:"device_is_mobile"`
	DeviceIsTablet              bool          `cql:"device_is_tablet"`
	DeviceIsPhone               bool          `cql:"device_is_phone"`
	LocationCityName            string        `cql:"location_city_name"`
	LocationCityID              uint          `cql:"location_city_id"`
	LocationCountryName         string        `cql:"location_country_name"`
	LocationCountryCode         string        `cql:"location_country_code"`
	LocationCountryID           uint          `cql:"location_country_id"`
	LocationContinentName       string        `cql:"location_continent_name"`
	LocationContinentCode       string        `cql:"location_continent_code"`
	LocationContinentID         uint          `cql:"location_continent_id"`
	LocationLatitude            float64       `cql:"location_latitude"`
	LocationLongitude           float64       `cql:"location_longitude"`
	LocationMetroCode           uint          `cql:"location_metro_code"`
	LocationTime                time.Time     `cql:"location_time"`
	LocationTimezone            string        `cql:"location_timezone"`
	LocationPostalCode          string        `cql:"location_postal_code"`
	LocationIsAnonymousProxy    bool          `cql:"location_is_anonymous_proxy"`
	LocationIsSatelliteProvider bool          `cql:"location_is_satellite_provider"`
	PageTitle                   string        `cql:"page_title"`
	PageHost                    string        `cql:"page_host"`
	PageURL                     string        `cql:"page_url"`
	ServerTime                  time.Time     `cql:"server_time"`
	ServerTimezone              time.Location `cql:"server_timezone"`
	ClientTime                  time.Time     `cql:"client_time"`
	ClientTimezone              time.Location `cql:"client_timezone"`
}
