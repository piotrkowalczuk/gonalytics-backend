package models

import (
	"labix.org/v2/mgo/bson"
)

type Report struct {
	Id                  bson.ObjectId `json: "id" bson:"_id,omitempty"`
	SiteId              int64         `json: "siteId" bson: "siteId"`
	Name                string        `json: "name" bson: "name"`
	AppName             string        `json: "appName" bson: "app_name"`
	Referrer            string        `json: "referrer" bson: "referrer"`
	Language            string        `json: "language" bson: "language"`
	Cookie              string        `json: "cookie" bson: "cookie"`
	UserAgent           string        `json: "userAgent" bson: "user_agent"`
	Java                bool          `json: "java" bson: "java"`
	BrowserVersion      string        `json: "browserVersions" bson: "browser_version"`
	BrowserVersionMinor string        `json: "browserVersionMinor" bson: "browser_version_minor"`
	ScreenWidth         int64         `json: "screenWidth" bson: "screen_width"`
	ScreenHeight        int64         `json: "screenHeight" bson: "screen_height"`
	WebsiteTitle        string        `json: "websiteTitle" bson: "website_title"`
	WebsiteHost         string        `json: "websiteHost" bson: "website_host"`
	WebsiteUrl          string        `json: "websiteUrl" bson: "website_url"`
}
