package fiscal

import (
	"fmt"
	"testing"
	"time"
)

func toDate(d time.Time) string {
	return fmt.Sprintf(`%04d-%02d-%02d`, d.Year(), d.Month(), d.Day())
}

func TestGetEndofYear(t *testing.T) {

	for i := 1998; i <= 2030; i++ {
		fc := NewYear(i)
		start := fc.YearStart

		end := fc.YearEnd

		days := float64(end.Sub(start).Hours()+24) / 24.
		weeks := days / 7
		t.Log(i, toDate(start), start.Weekday(), toDate(end), end.Weekday(), days, weeks, NumberOfWeeks(start, end))
	}
}

func TestChainSet(t *testing.T) {
	y := 2015
	//example of how to set other weekdays as beginning of week.
	fc := NewYear(y).SetStartWeeKDay(time.Sunday).SetTimeZone(time.UTC).SetStartEnd()
	start := fc.YearStart

	end := fc.YearEnd

	days := float64(end.Sub(start).Hours()+24) / 24.
	weeks := days / 7
	t.Log(y, toDate(start), start.Weekday(), toDate(end), end.Weekday(), days, weeks, NumberOfWeeks(start, end))
}
