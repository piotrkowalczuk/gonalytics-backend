package models

import (
	"time"
)

// AmountInTime ...
type AmountInTime struct {
	Amount   int64     `json:"amount" cql:"amount"`
	DateTime time.Time `json:"dateTime" cql:"date_time"`
}

// AmountInCountry ...
type AmountInCountry struct {
	Amount      int64  `json:"amount" cql:"amount"`
	CountryCode string `json:"countryCode" cql:"country_code"`
}
