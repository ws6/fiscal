package fiscal

import (
	"time"
)

//Calendar, given a Fiscal Year, generate quater, month, week and day
type Calendar struct {
	Quaters []*Quater

	_date  time.Time
	_start time.Time
	_end   time.Time

	_numWeeks        int
	_total_weeks     int
	_weeks_by_quater [][]int //13,13,13,13-14
	fc               *Year
}
type Quater struct {
	Months []*Month
}

type Month struct {
	Weeks []*Week
}

type Week struct {
	Days []*time.Time
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

func __sum(a []int) int {
	total := 0
	for _, n := range a {
		total += n
	}
	return total
}

func NewCalendar(fc *Year) *Calendar {
	ret := new(Calendar)
	ret.fc = fc
	ret._weeks_by_quater = [][]int{
		{4, 4, 5}, //month
		{4, 4, 5},
		{4, 4, 5},
		{4, 4, 5},
	} //default four quaters
	return ret.Generate()
}

func (self *Calendar) SetWeeksByQuater(w [][]int) *Calendar {
	if len(w) != 4 {
		return self
	}
	self._weeks_by_quater = w
	return self
}

func (self *Calendar) Generate() *Calendar {
	numWeeks := NumberOfWeeks(self.fc.YearStart, self.fc.YearEnd)
	if numWeeks == 53 {
		self._weeks_by_quater[len(self._weeks_by_quater)-1][len(self._weeks_by_quater[len(self._weeks_by_quater)-1])-1] += 1
	}
	days := numWeeks * 7
	return self
}
