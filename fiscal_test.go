package fiscal

import (
	"fmt"
	"strings"
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

func TestCalender(t *testing.T) {
	year := 2019

	cal := NewCal(year)

	for _, q := range cal.Quaters {
		t.Log(`Q `, q.OrderRel)
		for _, m := range q.Months {
			t.Log(time.Month(m.OrderAbs))

			for _, w := range m.Weeks {
				//				t.Log(`week `, w.OrderRel, w.OrderAbs)
				weekdays := []string{}
				for _, d := range w.Days {
					weekdays = append(weekdays, toDate(*d))
				}

				t.Log(`W `, w.OrderRel, w.OrderAbs, strings.Join(weekdays, " "))
			}
		}

	}
}

func sameFiscalDate(d1, d2 Date) bool {

	return d1.Year == d2.Year &&
		d1.Quater == d2.Quater &&
		d1.Month == d2.Month &&
		d1.Week == d2.Week &&
		d1.WeekOfYear == d2.WeekOfYear
}

func TestToDate(t *testing.T) {
	//year int, month Month, day, hour, min, sec, nsec int, loc *Location
	dates := []time.Time{
		time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2018, 12, 31, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 4, 29, 0, 0, 0, 0, time.UTC),
		time.Date(2019, 4, 28, 0, 0, 0, 0, time.UTC),
	}
	results := []Date{

		{2020, 1, 1, 1, 1},
		{2019, 1, 1, 1, 1},
		{2019, 2, 5, 1, 18},
		{2019, 2, 4, 4, 17},
	}
	for i, d := range dates {
		fisDate := ToDate(d)
		if !sameFiscalDate(*fisDate, results[i]) {
			t.Log(`fiscal date ->`, fisDate, d.Weekday())
			t.Fatal(`not correct`)
		}

	}

}
