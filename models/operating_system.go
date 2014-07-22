package models

type OperatingSystem struct {
	Name    string `json:"name" bson:"name"`
	Version string `json:"version" bson:"version"`
}
