package jalali

import (
	"time"
)

// 1461 = 365*4 + 4/4
// 36525 = 365*100 + 100/4
// 146097 = 365*400 + 400/4 - 400/100 + 400/400
var breaks = [...]int{-61, 9, 38, 199, 426, 686, 756, 818, 1111, 1181, 1210,
	1635, 2060, 2097, 2192, 2262, 2324, 2394, 2456, 3178}

type Calendar struct {
	year    int
	month   int
	day     int
	hour    int
	min     int
	sec     int
	weekDay time.Weekday
	loc     *time.Location
}

// ToJalali converts Gregorian to Jalali date. Error is not nil if Jalali
// year() passed to function is not valid.
func ToJalali(gy, gm, gd int) (Calendar, error) {
	c, err := d2j(g2d(gy, gm, gd))
	return c, err
}

//func ToJalali(gregorianYear int, gregorianMonth time.month(), gregorianDay int) (int, month(), int, error) {
//	jy, jm, jd, err := d2j(g2d(gregorianYear, int(gregorianMonth), gregorianDay))
//	return jy, month()(jm), jd, err
//}

// ToGregorian converts Jalali to Gregorian date. Error is not nil if Jalali
// year() passed to function is not valid.
func (c Calendar) ToGregorian() (Calendar, error) {
	jdn, err := j2d(c.GetYear(), c.GetMonth(), c.GetDay())
	if err != nil {
		return New(), err
	}

	georgianCall := d2g(jdn)
	return georgianCall, nil
}


func jalali(jy int) (int, int, int, error) {
	bl, gy, leapJ, jp := len(breaks), jy+621, -14, breaks[0]
	jump := 0

	//check for years is valid or not
	if jy < jp || jy >= breaks[bl-1] {
		return 0, 0, 0, &ErrorInvalidYear{jy}
	}

	// Find the limiting years for the Jalali year() jy.
	for i := 1; i < bl; i++ {
		jm := breaks[i]
		jump := jm - jp
		if jy < jm {
			break
		}
		// j_day_no = nt(jy//33)*8 + (jy%33+3)//4
		leapJ += div(jump, 33)*8 + div(mod(jump, 33), 4)
		jp = jm
	}
	n := jy - jp

	// Find the number of leap years from AD 621 to the beginning
	// of the current Jalali year() in the Persian JalaliCal.
	leapJ += div(n, 33)*8 + div(mod(n, 33)+3, 4)
	if mod(jump, 33) == 4 && jump-n == 4 {
		leapJ++
	}

	// And the same in the Gregorian JalaliCal (until the year() gy).
	leapG := div(gy, 4) - div((div(gy, 100)+1)*3, 4) - 150

	// Determine the Gregorian date of Farvardin the 1st.
	march := 20 + leapJ - leapG

	// Find how many years have passed since the last leap year().
	if jump-n < 6 {
		n -= jump + div(jump+4, 33)*33
	}
	leap := mod(mod(n+1, 33)-1, 4)
	if leap == -1 {
		leap = 4
	}

	return leap, gy, march, nil
}

func (c Calendar) GetYear() int {
	return c.year
}

func (c Calendar) GetDay() int {
	return c.day
}
func (c Calendar) GetMonth() int {
	return c.month
}

func (c Calendar) GetHour() int {
	return c.hour
}

func (c Calendar) GetMinute() int {
	return c.min
}
func (c Calendar) GetSecond() int {
	return c.sec
}

func (c Calendar) GetWeekday() time.Weekday {
	return c.weekDay
}

func (c Calendar) GetLocation() *time.Location {
	return c.loc
}

func (c Calendar) SetYear(y int) Calendar {
	c.year = y
	return c
}

func (c Calendar) SetDay(d int) Calendar {
	c.day = d
	return c
}
func (c Calendar) SetMonth(m int) Calendar {
	c.month = m
	return c
}

func (c Calendar) SetHour(h int) Calendar {
	c.hour = h
	return c
}

func (c Calendar) SetMinute(m int) Calendar {
	c.min = m
	return c
}
func (c Calendar) SetSecond(s int) Calendar {
	c.sec = s
	return c
}

func (c Calendar) SetWeekday(wd time.Weekday) Calendar {
	c.weekDay = wd
	return c
}

func (c Calendar) SetLocation(l *time.Location) Calendar {
	c.loc = l
	return c
}

// New return empty calendar
func New() Calendar {
	return Calendar{}
}

// Date return arbitrary date time
func Date(year, month, day, hour, min, sec int) Calendar {
	c := Calendar{}
	return c.
		SetYear(year).
		SetMonth(month).
		SetDay(day).
		SetHour(hour).
		SetMinute(min).
		SetSecond(sec)

}

// Iran returns a pointer to time.Location of Asia/Tehran
func Iran() *time.Location {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		loc = time.FixedZone("Asia/Tehran", 12600) // UTC + 03:30
	}
	return loc
}

// Unix convert Unix Time to Jalali Date
func Unix(unixDate int64) (Calendar, error) {
	unixTimeUTC := time.Unix(unixDate, 0) //gives unix time stamp in utc
	c, err := d2j(g2d(unixTimeUTC.Year(), int(unixTimeUTC.Month()), unixTimeUTC.Day()))
	return c, err
}

// Now get system time as unix form and return it to persian calendar
func Now() Calendar {
	c, _ := Unix(time.Now().Unix())
	return c.SetHour(time.Now().Hour()).
		SetMinute(time.Now().Minute()).
		SetSecond(time.Now().Second())
}

// JTU convert persian calendar to unix format
func JTU(cal Calendar) int64 {
	jt, _ := cal.ToGregorian()
	gt := time.Date(jt.GetYear(), time.Month(jt.GetMonth()), jt.GetDay(), 12, 15, 60, 0, time.UTC)
	return gt.Unix()
}

// JDiff calculate difference between two jalali time
// in function determine what date is exceeded another date
func (c Calendar) JDiff(cal Calendar) int64 {
	var e, b int64
	if c.GetYear() >= cal.GetYear() && c.GetMonth() >= cal.GetMonth() && c.GetDay() >= cal.GetDay() {
		b = JTU(cal)
		e = JTU(c)
	} else {
		b = JTU(c)
		e = JTU(cal)
	}
	return (e - b) / 86400
}

// MonthEnName return persian month name in english lang
func (c Calendar) MonthEnName() string {
	return toEng(c.month)
}

// MonthPersianName return persian month name in farsi lang
func (c Calendar) MonthPersianName() string {
	return toPersian(c.month)
}
