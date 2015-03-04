package models

// MetricCounterEntity ...
type MetricCounterEntity struct {
	DimensionsNames  string `json:"dimensionsNames"`
	DimensionsValues string `json:"dimensionsValues"`
	Count            int64  `json:"count"`
}

// MetricDayCounterEntity ...
type MetricDayCounterEntity struct {
	MetricCounterEntity
	MadeAtDay   int `json:"madeAtDay" cql:"made_at_day"`
	MadeAtMonth int `json:"madeAtMonth" cql:"made_at_month"`
	MadeAtYear  int `json:"madeAtYear" cql:"made_at_year"`
}

// MetricMonthCounterEntity ...
type MetricMonthCounterEntity struct {
	MetricCounterEntity
	MadeAtDay  int `json:"madeAtDay" cql:"made_at_day"`
	MadeAtYear int `json:"madeAtYear" cql:"made_at_month"`
}

// MetricYearCounterEntity ...
type MetricYearCounterEntity struct {
	MetricCounterEntity
	MadeAtDay int `json:"madeAtDay"`
}
