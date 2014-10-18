package models

type OperatingSystem struct {
	Name    string `json:"name" cql:"name"`
	Version string `json:"version" cql:"version"`
}
