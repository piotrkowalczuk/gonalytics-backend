package models

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
