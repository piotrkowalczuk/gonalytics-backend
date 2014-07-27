package models

import (
    "time"
)

type AmountInTime struct {
    Amount int64 `json:"amount" bson:"amunt"`
    DateTime time.Time  `json:"dateTime" bson:"date_time"`
}
