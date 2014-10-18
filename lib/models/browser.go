package models

// Browser represents browser information.
type Browser struct {
	Name         string  `json:"name" cql:"name"`
	Version      string  `json:"version" cql:"version"`
	MajorVersion string  `json:"majorVersion" cql:"major_version"`
	UserAgent    string  `json:"userAgent" cql:"user_agent"`
	Platform     string  `json:"platform" cql:"platform"`
	Cookie       bool    `json:"cookie" cql:"cookie"`
	Plugins      Plugins `json:"plugins" cql:"plugins"`
	IsOnline     bool    `json:"isOnline" cql:"is_online"`
	Window       Window  `json:"window" cql:"window"`
}
