package fiscal

import (
	"math"
	"time"
)

//fiscal ->445 or 544 fiscal calendar calculater

var (
	//default for user to set
	DEFAULT_TIME_ZOME      *time.Location
	DEFAULT_START_WEEK_DAY = time.Monday
	DEFAULT_START_MONTH    = time.January
)

type Year struct {
	START_WEEK_DAY time.Weekday
	START_MONTH    time.Month
	Zone           *time.Location

	YearStart time.Time
	YearEnd   time.Time
	_year     int
}

func NewYear(year int) *Year {

	ret := new(Year)
	ret._year = year
	//Lazy Default
	ret.Zone = DEFAULT_TIME_ZOME
	ret.START_MONTH = DEFAULT_START_MONTH
	if DEFAULT_TIME_ZOME == nil {
		ret.Zone, _ = time.LoadLocation("America/Los_Angeles")
	}

	ret.START_WEEK_DAY = DEFAULT_START_WEEK_DAY

	return ret.SetStartEnd()
}

func (self *Year) SetStartWeeKDay(wd time.Weekday) *Year {
	self.START_WEEK_DAY = wd
	return self
}
func (self *Year) SetStartMonth(m time.Month) *Year {
	if m > 0 && m <= 12 {
		self.START_MONTH = m
	}
	return self
}
func (self *Year) SetTimeZone(z *time.Location) *Year {
	self.Zone = z
	return self
}

func (self *Year) SetStartEnd() *Year {
	self.YearStart = self.GetStartOfYear(self._year)
	self.YearEnd = self.GetEndOfYear(self._year)
	return self
}

func (self *Year) GetStartOfYear(year int) time.Time {

	//if leap year, align with next year's start must be 52 weeks
	if IsLeapYear(year) {
		nextYearStart := self.GetStartOfYear(year + 1)
		return nextYearStart.AddDate(0, 0, -7*52)
	}

	firstDay := time.Date(year, self.START_MONTH, 1, 0, 0, 0, 0, self.Zone)
	return firstDay.AddDate(0, 0, int(self.START_WEEK_DAY-firstDay.Weekday()))
}

func (self *Year) GetEndOfYear(year int) time.Time {
	return self.GetStartOfYear(year+1).AddDate(0, 0, -1)
}

func NumberOfWeeks(start, end time.Time) int {

	days := end.Sub(start).Hours() / 24
	weeks := days / 7

	return int(math.Ceil(float64(weeks)))

}

func IsLeapYear(y int) bool {
	year := time.Date(y, time.December, 31, 0, 0, 0, 0, time.Local)
	return year.YearDay() > 365
}
