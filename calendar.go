package fiscal

import (
	"time"
)

//allow user to reset this
var DEFAULT_WEEKS_BY_QUATER = [][]int{
	{4, 4, 5}, //month; it can be 445 or 544; user maintain the correctness and consistency.
	{4, 4, 5},
	{4, 4, 5},
	{4, 4, 5},
} //default four quaters
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
	OrderRel int

	Months []*Month
}

type Month struct {
	OrderRel      int
	OrderAbs      int
	_weeksInMonth int
	Weeks         []*Week
}

type Week struct {
	OrderRel int
	OrderAbs int
	Days     []*time.Time
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
	ret._weeks_by_quater = DEFAULT_WEEKS_BY_QUATER
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
	//	days := numWeeks * 7
	monthOrder := 0
	weekOrder := 0
	dayOrder := 0
	for qi, _quater := range self._weeks_by_quater {

		quater := new(Quater)
		quater.OrderRel = qi + 1
		self.Quaters = append(self.Quaters, quater)

		for mi, weeksInMonth := range _quater {

			month := new(Month)
			month._weeksInMonth = weeksInMonth
			month.OrderRel = mi + 1
			monthOrder += 1
			month.OrderAbs = monthOrder
			quater.Months = append(quater.Months, month)

			//create weeks
			for wi := 0; wi < weeksInMonth; wi++ {
				week := new(Week)
				week.OrderRel = wi + 1
				weekOrder += 1
				week.OrderAbs = weekOrder

				month.Weeks = append(month.Weeks, week)
				for di := 0; di < 7; di++ {
					day := self.fc.YearStart.AddDate(0, 0, dayOrder)
					dayOrder += 1
					week.Days = append(week.Days, &day)
				}
			}
		}

	}
	return self
}

func NewCal(year int) *Calendar {
	return NewCalendar(NewYear(year))
}
