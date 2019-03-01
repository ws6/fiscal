package fiscal

//day.go provie API for day to fiscal day and vise versa
import (
	"time"
)

type Date struct {
	Year       int
	Quater     int
	Month      int
	Week       int
	WeekOfYear int
}

func DateEqual(a, b time.Time) bool {
	return dateToInt(a) == dateToInt(b)
}

func dateToInt(d time.Time) int {
	return d.Year()*10000 + int(d.Month()*100) + d.Day()
}

//is a before b
func DateLessThan(a, b time.Time) bool {
	return dateToInt(a) < dateToInt(b)

}
func DateGreaterThan(a, b time.Time) bool {
	return dateToInt(a) > dateToInt(b)
}

func IsInBetween(q, start, end time.Time) bool {
	if DateEqual(q, start) || DateEqual(q, end) {
		return true
	}

	return DateGreaterThan(q, start) && DateLessThan(q, end)
}

func (self *Calendar) Which(d time.Time) *Date {
	for _, q := range self.Quaters {
		for _, m := range q.Months {

			for _, w := range m.Weeks {
				for _, _d := range w.Days {
					if DateEqual(*_d, d) {

						ret := new(Date)
						ret.Year = self.fc._year
						ret.Quater = q.OrderRel
						ret.Month = m.OrderAbs
						ret.Week = w.OrderRel
						ret.WeekOfYear = w.OrderAbs

						return ret

					}
				}
			}
		}
	}
	return nil
}

//!!!everything takes default except caller change the global default before calling
func ToDate(d time.Time) *Date {

	cal := NewCal(d.Year())

	//if in this fiscal year?
	if !IsInBetween(d, cal.fc.YearStart, cal.fc.YearEnd) {
		//look it up then report it
		if DateLessThan(d, cal.fc.YearStart) {
			cal = NewCal(d.Year() - 1)
		}
		if DateGreaterThan(d, cal.fc.YearStart) {
			cal = NewCal(d.Year() + 1)
		}
	}

	return cal.Which(d)
}
