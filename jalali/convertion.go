package jalali

import (
	"time"
)

// 1461 = 365*4 + 4/4
// 36525 = 365*100 + 100/4
// 146097 = 365*400 + 400/4 - 400/100 + 400/400
var breaks = [...]int{-61, 9, 38, 199, 426, 686, 756, 818, 1111, 1181, 1210,
	1635, 2060, 2097, 2192, 2262, 2324, 2394, 2456, 3178}

type PersianCalendar struct {
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
func ToJalali(gy, gm, gd int) (PersianCalendar, error) {
	c, err := d2j(g2d(gy, gm, gd))
	return c, err
}

//func ToJalali(gregorianYear int, gregorianMonth time.month(), gregorianDay int) (int, month(), int, error) {
//	jy, jm, jd, err := d2j(g2d(gregorianYear, int(gregorianMonth), gregorianDay))
//	return jy, month()(jm), jd, err
//}

// ToGregorian converts Jalali to Gregorian date. Error is not nil if Jalali
// year() passed to function is not valid.
func (c PersianCalendar) ToGregorian() (PersianCalendar, error) {
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

func (c PersianCalendar) GetYear() int {
	return c.year
}

func (c PersianCalendar) GetDay() int {
	return c.day
}
func (c PersianCalendar) GetMonth() int {
	return c.month
}

func (c PersianCalendar) GetHour() int {
	return c.hour
}

func (c PersianCalendar) GetMinute() int {
	return c.min
}
func (c PersianCalendar) GetSecond() int {
	return c.sec
}

func (c PersianCalendar) GetWeekday() time.Weekday {
	return c.weekDay
}

func (c PersianCalendar) GetLocation() *time.Location {
	return c.loc
}

func (c PersianCalendar) SetYear(y int) PersianCalendar {
	c.year = y
	return c
}

func (c PersianCalendar) SetDay(d int) PersianCalendar {
	c.day = d
	return c
}
func (c PersianCalendar) SetMonth(m int) PersianCalendar {
	c.month = m
	return c
}

func (c PersianCalendar) SetHour(h int) PersianCalendar {
	c.hour = h
	return c
}

func (c PersianCalendar) SetMinute(m int) PersianCalendar {
	c.min = m
	return c
}
func (c PersianCalendar) SetSecond(s int) PersianCalendar {
	c.sec = s
	return c
}

func (c PersianCalendar) SetWeekday(wd time.Weekday) PersianCalendar {
	c.weekDay = wd
	return c
}

func (c PersianCalendar) SetLocation(l *time.Location) PersianCalendar {
	c.loc = l
	return c
}

// New return empty calendar
func New() PersianCalendar {
	return PersianCalendar{}
}

// Date return arbitrary date time
func Date(year, month, day, hour, min, sec int) PersianCalendar {
	c := PersianCalendar{}
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
func Unix(unixDate int64) (PersianCalendar, error) {
	unixTimeUTC := time.Unix(unixDate, 0) //gives unix time stamp in utc
	c, err := d2j(g2d(unixTimeUTC.Year(), int(unixTimeUTC.Month()), unixTimeUTC.Day()))
	return c, err
}

// Now get system time as unix form and return it to persian calendar
func Now() PersianCalendar {
	c, _ := Unix(time.Now().Unix())
	return c.SetHour(time.Now().Hour()).
		SetMinute(time.Now().Minute()).
		SetSecond(time.Now().Second())
}

// JTU convert persian calendar to unix format
func JTU(cal PersianCalendar) int64 {
	jt, _ := cal.ToGregorian()
	gt := time.Date(jt.GetYear(), time.Month(jt.GetMonth()), jt.GetDay(), 12, 15, 60, 0, time.UTC)
	return gt.Unix()
}

// JDiff calculate difference between two jalali time
// in function determine what date is exceeded another date
func (c PersianCalendar) JDiff(cal PersianCalendar) int64 {
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
func (c PersianCalendar) MonthEnName() string {
	return toEng(c.month)
}

// MonthPersianName return persian month name in farsi lang
func (c PersianCalendar) MonthPersianName() string {
	return toPersian(c.month)
}
