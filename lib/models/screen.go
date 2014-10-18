package models

type Screen struct {
	Width  int64 `json:"width" cql:"width"`
	Height int64 `json:"height" cql:"height"`
}
