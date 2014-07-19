package models

type Action struct {
	Id              string           `json:"id" bson:"_id,omitempty"`
	UserId          string           `json:"userId" bson:"userId"`
	SiteId          int64            `json:"siteId" bson:"siteId"`
	Referrer        string           `json:"referrer" bson:"referrer"`
	Language        string           `json:"language" bson:"language"`
	Browser         *Browser         `json:"browser" bson:"browser"`
	Screen          *Screen          `json:"screen" bson:"screen"`
	Website         *Website         `json:"website" bson:"website"`
	OperatingSystem *OperatingSystem `json:"os" bson:"os"`
	Device          *Device          `json:"device" bson:"device"`
	CreatedAt       *MongoDate       `json:"createdAt" bson:"created_at"`
}

type Website struct {
	Title string `json:"title" bson:"title"`
	Host  string `json:"host" bson:"host"`
	Url   string `json:"url" bson:"url"`
}

type Screen struct {
	Width  int64 `json:"width" bson:"width"`
	Height int64 `json:"height" bson:"height"`
}

type Window struct {
	Width  int64 `json:"width" bson:"width"`
	Height int64 `json:"height" bson:"height"`
}

type Browser struct {
	Name         string  `json:"name" bson:"name"`
	Version      string  `json:"version" bson:"version"`
	MajorVersion string  `json:"majorVersion" bson:"major_version"`
	UserAgent    string  `json:"userAgent" bson:"user_agent"`
	Platform     string  `json:"platform" bson:"platform"`
	Cookie       bool    `json:"cookie" bson:"cookie"`
	Plugins      Plugins `json:"plugins" bson:"plugins"`
	IsOnline     bool    `json:"isOnline" bson:"is_online"`
	Window       Window  `json:"window" bson:"window"`
}

type Plugins struct {
	Java bool `json:"java" bson:"java"`
}

type OperatingSystem struct {
	Name string `json:"name" bson:"name"`
}

type Device struct {
	Name     string `json:"name" bson:"name"`
	IsMobile bool   `json:"isMobile" bson:"is_mobile"`
	IsTablet bool   `json:"isTablet" bson:"is_tablet"`
	IsPhone  bool   `json:"isPhone" bson:"is_phone"`
}
