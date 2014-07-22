package models

type Screen struct {
	Width  int64 `json:"width" bson:"width"`
	Height int64 `json:"height" bson:"height"`
}
