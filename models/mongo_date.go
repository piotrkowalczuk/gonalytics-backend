package models

import (
	"strconv"
	"time"
)

type MongoDate struct {
	DateTime time.Time `json:"dateTime" bson:"date_time"`
	Bucket   []string  `json:"bucket" bson:"bucket"`
}

func NewMongoDate(dateTime time.Time) *MongoDate {
	dtb := new(MongoDate)

	dtb.DateTime = dateTime
	dtb.AddSecond(dateTime)
	dtb.AddMinute(dateTime)
	dtb.AddHour(dateTime)
	dtb.AddDay(dateTime)
	dtb.AddWeek(dateTime)
	dtb.AddMonth(dateTime)
	dtb.AddYear(dateTime)

	return dtb
}

func (dtb *MongoDate) AppendBucket(value string) {
	dtb.Bucket = append(dtb.Bucket, value)
}

func (dtb *MongoDate) AddSecond(t time.Time) {
	second := t.UTC().Format("2006-01-02 15:04:05") + "-second"

	dtb.AppendBucket(second)
}

func (dtb *MongoDate) AddMinute(t time.Time) {
	minute := t.UTC().Format("2006-01-02 15:04") + "-minute"

	dtb.AppendBucket(minute)
}

func (dtb *MongoDate) AddHour(t time.Time) {
	hour := t.UTC().Format("2006-01-02 15") + "-hour"

	dtb.AppendBucket(hour)
}

func (dtb *MongoDate) AddDay(t time.Time) {
	day := t.UTC().Format("2006-01-02") + "-day"

	dtb.AppendBucket(day)
}

func (dtb *MongoDate) AddWeek(t time.Time) {
	year, week := t.UTC().ISOWeek()

	dtb.AppendBucket(strconv.FormatInt(int64(year), 10) + "-" + strconv.FormatInt(int64(week), 10) + "-week")
}

func (dtb *MongoDate) AddMonth(t time.Time) {
	month := t.UTC().Format("2006-01") + "-month"

	dtb.AppendBucket(month)
}

func (dtb *MongoDate) AddYear(t time.Time) {
	year := t.UTC().Format("2006") + "-year"

	dtb.AppendBucket(year)
}
