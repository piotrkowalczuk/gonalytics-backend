package models

import (
	"testing"
	"time"
)

const pattern = "Jan 2, 2006 at 3:04pm (MST)"

func NewVisitAverageDurationMock(firstActionAtString, lastActionAtString string) *Visit {
	firstActionAt, _ := time.Parse(pattern, firstActionAtString)
	lastActionAt, _ := time.Parse(pattern, lastActionAtString)

	visit := Visit{
		FirstActionAt: firstActionAt,
		LastActionAt:  lastActionAt,
	}

	return &visit
}

func TestAverageDuration1(t *testing.T) {
	dr1 := NewVisitAverageDurationMock(
		"Feb 3, 2013 at 7:14pm (PST)",
		"Feb 3, 2013 at 7:24pm (PST)",
	)

	dr2 := NewVisitAverageDurationMock(
		"Feb 3, 2013 at 7:14pm (PST)",
		"Feb 3, 2013 at 7:34pm (PST)",
	)

	dr3 := NewVisitAverageDurationMock(
		"Feb 3, 2013 at 7:14pm (PST)",
		"Feb 3, 2013 at 7:44pm (PST)",
	)

	dr4 := NewVisitAverageDurationMock(
		"Feb 3, 2013 at 7:14pm (PST)",
		"Feb 3, 2013 at 7:54pm (PST)",
	)

	visits := []*Visit{dr1, dr2, dr3, dr4}

	testVisitAvarageDuration(t, 25, visits)
}

func TestAverageDuration2(t *testing.T) {
	dr1 := NewVisitAverageDurationMock(
		"Feb 3, 2013 at 7:14pm (PST)",
		"Feb 3, 2013 at 7:15pm (PST)",
	)

	dr2 := NewVisitAverageDurationMock(
		"Feb 3, 2013 at 7:14pm (PST)",
		"Feb 3, 2013 at 7:15pm (PST)",
	)

	dr3 := NewVisitAverageDurationMock(
		"Feb 3, 2013 at 7:14pm (PST)",
		"Feb 3, 2013 at 7:15pm (PST)",
	)

	dr4 := NewVisitAverageDurationMock(
		"Feb 3, 2013 at 7:12pm (PST)",
		"Feb 3, 2013 at 7:15pm (PST)",
	)

	visits := []*Visit{dr1, dr2, dr3, dr4}

	testVisitAvarageDuration(t, 1.5, visits)
}

func TestAverageDuration3(t *testing.T) {
	visits := []*Visit{}

	testVisitAvarageDuration(t, 1.5, visits)
}

func testVisitAvarageDuration(t *testing.T, expectedDuration float64, visits []*Visit) {
	ad := VisitsAverageDuration(visits)

	if ad.Minutes() != float64(expectedDuration) {
		t.Errorf("Average duration should be equal to %d but is %#v", expectedDuration, ad.Minutes())
	}
}
