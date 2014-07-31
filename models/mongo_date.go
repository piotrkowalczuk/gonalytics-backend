package models

import (
	"strconv"
	"time"
)

// MongoDate ...
type MongoDate struct {
	DateTime *time.Time `json:"dateTime" bson:"date_time"`
	Bucket   []string   `json:"bucket" bson:"bucket"`
}

// NewMongoDate ...
func NewMongoDate(dateTime *time.Time) *MongoDate {
	dtb := new(MongoDate)

	dtb.DateTime = dateTime
	dtb.addSecond(dateTime)
	dtb.addMinute(dateTime)
	dtb.addHour(dateTime)
	dtb.addDay(dateTime)
	dtb.addWeek(dateTime)
	dtb.addMonth(dateTime)
	dtb.addYear(dateTime)

	return dtb
}

// AppendBucket ...
func (dtb *MongoDate) AppendBucket(value string) {
	dtb.Bucket = append(dtb.Bucket, value)
}

func (dtb *MongoDate) addSecond(t *time.Time) {
	second := t.UTC().Format("2006-01-02 15:04:05") + "-second"

	dtb.AppendBucket(second)
}

func (dtb *MongoDate) addMinute(t *time.Time) {
	minute := t.UTC().Format("2006-01-02 15:04") + "-minute"

	dtb.AppendBucket(minute)
}

func (dtb *MongoDate) addHour(t *time.Time) {
	hour := t.UTC().Format("2006-01-02 15") + "-hour"

	dtb.AppendBucket(hour)
}

func (dtb *MongoDate) addDay(t *time.Time) {
	day := t.UTC().Format("2006-01-02") + "-day"

	dtb.AppendBucket(day)
}

func (dtb *MongoDate) addWeek(t *time.Time) {
	year, week := t.UTC().ISOWeek()

	dtb.AppendBucket(strconv.FormatInt(int64(year), 10) + "-" + strconv.FormatInt(int64(week), 10) + "-week")
}

func (dtb *MongoDate) addMonth(t *time.Time) {
	month := t.UTC().Format("2006-01") + "-month"

	dtb.AppendBucket(month)
}

func (dtb *MongoDate) addYear(t *time.Time) {
	year := t.UTC().Format("2006") + "-year"

	dtb.AppendBucket(year)
}
