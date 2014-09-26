package models

import (
	"time"
)

// AmountInTime ...
type AmountInTime struct {
	Amount   int64     `json:"amount" bson:"amount"`
	DateTime time.Time `json:"dateTime" bson:"date_time"`
}

// AmountInCountry ...
type AmountInCountry struct {
	Amount      int64  `json:"amount" bson:"amount"`
	CountryCode string `json:"countryCode" bson:"country_code"`
}
