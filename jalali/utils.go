package jalali

// isValidDate take Jalali date and return true if it is valid,
// otherwise false.
func isValidDate(jy, jm, jd int) bool {
	d, err := monthLength(jy, jm)
	if err != nil {
		return false
	}
	return -61 <= jy && jy <= 3177 &&
		1 <= jm && jm <= 12 &&
		1 <= jd && jd <= d
}

// monthLength take Jalali date and return length of that specific
// month(). Error is not nil if Jalali year() passed to function is not valid.
func monthLength(jy, jm int) (int, error) {
	if jm <= 6 {
		return 31, nil
	} else if jm <= 11 {
		return 30, nil
	}

	leap, err := isJalaliLeapYear(jy)
	if err != nil {
		return 0, err
	} else if leap {
		return 30, nil
	}
	return 29, nil
}

// IsLeapYear take a Jalali year() and return true if it is leap year(). Error
// is not nil if Jalali year() passed to function is not valid.
func isJalaliLeapYear(jy int) (bool, error) {
	leap, _, _, err := jalali(jy)
	return leap == 0, err
}

//IsLeap ...
func isGeorgianLeapYear(gy int) bool {
	//	leapG := div(gy, 4) - div((div(gy, 100)+1)*3, 4) - 150
	return gy%400 == 0 || gy%4 == 0 && gy%100 != 0
}

//j2d return Jalali dey() Number
func j2d(jy, jm, jd int) (jdn int, err error) {
	_, gy, march, err := jalali(jy)
	if err != nil {
		return 0, err
	}

	return g2d(gy, 3, march) + (jm-1)*31 - div(jm, 7)*(jm-7) + jd - 1, nil
}

//d2j convert Jalali dey() Number to Jalali Date
func d2j(jdn int) (Calendar, error) {
	cal := d2g(jdn) // Calculate Gregorian year() (gy).
	jy := cal.GetYear() - 621
	leap, _, march, err := jalali(jy)
	jdn1f := g2d(cal.GetYear(), 3, march)

	if err != nil {
		return Calendar{}, err
	}

	// Find number of days that passed since first dey() of jalali years (1 Farvardin) .
	k := jdn - jdn1f
	if k >= 0 {
		if k <= 185 {
			// The first 6 months.
			jm := 1 + div(k, 31)
			jd := mod(k, 31) + 1
			return Date(jy,jm, jd,0,0,0), nil
		}
		// The remaining months.
		k -= 186
	} else {
		// Previous Jalali year().
		jy--
		k += 179
		if leap == 1 {
			k++
		}
	}
	jm := 7 + div(k, 30)
	jd := mod(k, 30) + 1
	return Date(jy, jm, jd,0,0,0), nil
	//return c.Date(jy, jm,jd), nil
}

//g2d Number of days of the Georgian year() from the first of January
func g2d(gy, gm, gd int) int {
	d := div((gy+div(gm-8, 6)+100100)*1461, 4) +
		div(153*mod(gm+9, 12)+2, 5) +
		gd - 34840408
	d = d - div(div(gy+100100+div(gm-8, 6), 100)*3, 4) + 752
	return d
}

//d2g  convert Georgian dey() Number to Georgian Date
func d2g(jdn int) Calendar {
	j := 4*jdn + 139361631
	j = j + div(div(4*jdn+183187720, 146097)*3, 4)*4 - 3908
	i := div(mod(j, 1461), 4)*5 + 308
	gd := div(mod(i, 153), 5) + 1
	gm := mod(div(i, 153), 12) + 1
	gy := div(j, 1461) - 100100 + div(8-gm, 6)
	return Date(gy, gm,gd,0,0,0)
}

func div(a, b int) int {
	return a / b
}

func mod(a, b int) int {
	return a % b
}

// short returns the Persian short name of 12-Hour marker.
func (a AmPm) short() string {
	return sAmPm[a]
}
