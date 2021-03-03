package jalali

import (
	"fmt"
	"strconv"
	"strings"
)

//func (c PersianCalendar) Format(format string) string {
//	r := strings.NewReplacer(
//		"yyyy", strconv.Itoa(c.GetYear()),
//		//"yyy", strconv.Itoa(c.GetYear()),
//		//"yy", strconv.Itoa(c.GetYear())[2:],
//		//"y", strconv.Itoa(c.GetYear())[3:],
//		"YYYY", enToFa.Replace(strconv.Itoa(c.GetYear())),// ۱۳۹۹
//		//"YYY", enToFa.Replace(strconv.Itoa(c.GetYear())),
//		//"YY", enToFa.Replace(strconv.Itoa(c.GetYear())[2:]),// 99
//		//"Y", enToFa.Replace(strconv.Itoa(c.GetYear()))[3:],//9
//		"mm", fmt.Sprintf("%02d", c.GetMonth()),
//		//"m", strconv.Itoa(c.GetMonth()),
//		"dd", fmt.Sprintf("%02d", c.GetDay()),
//		"MM", toEng(c.GetMonth()),
//		//"M", toPersian(c.GetMonth()),
//		"DD", toPersian(c.GetDay()),
//		//"H",strconv.Itoa(c.GetHour()),
//		"HH", fmt.Sprintf("%02d", c.GetHour()),
//		//"h",enToFa.Replace(strconv.Itoa(c.GetHour())),
//		"hh",enToFa.Replace(enToFa.Replace(strconv.Itoa(c.GetHour()))),
//		//"T",strconv.Itoa(c.GetMinute()),
//		"TT", fmt.Sprintf("%02d", c.GetMinute()),
//		//"t",enToFa.Replace(strconv.Itoa(c.GetMinute())),
//		"tt",enToFa.Replace(fmt.Sprintf("%02d", c.GetMinute())),
//		//"S",strconv.Itoa(c.GetSecond()),
//		"SS", fmt.Sprintf("%02d", c.GetSecond()),
//		//"s",enToFa.Replace(strconv.Itoa(c.GetSecond())),
//		"ss",enToFa.Replace(fmt.Sprintf("%02d", c.GetSecond())),
//	)
//	return r.Replace(format)
//}

func (c PersianCalendar) Format(format string) string {
	r := strings.NewReplacer(
		// Eng Format ......
		// Year
		"{yyyy}", strconv.Itoa(c.GetYear()),
		"{yyy}", strconv.Itoa(c.GetYear())[1:],
		"{yy}", strconv.Itoa(c.GetYear())[2:],
		"{y}", strconv.Itoa(c.GetYear())[3:],
		// Month
		"{mm}", fmt.Sprintf("%02d", c.GetMonth()),
		"{m}", fmt.Sprintf("%d", c.GetMonth()),
		"{mo}", toEng(c.GetMonth()),
		// Day
		"{dd}", fmt.Sprintf("%02d", c.GetDay()),
		"{d}", fmt.Sprintf("%d", c.GetDay()),
		"{wd}", wdToEn.Replace(strconv.Itoa(int(c.GetWeekday()))),
		// Hour
		"{hh}", fmt.Sprintf("%02d", c.GetHour()),
		"{h}", fmt.Sprintf("%d", c.GetHour()),
		// Minute
		"{tt}",fmt.Sprintf("%02d", c.GetMinute()),
		"{t}",fmt.Sprintf("%d", c.GetMinute()),
		// Second
		"{ss}",fmt.Sprintf("%02d", c.GetSecond()),
		"{s}",fmt.Sprintf("%d", c.GetSecond()),
		// Persian format .............
		// Year
		"{YYYY}", enToFa.Replace(strconv.Itoa(c.GetYear())),// ۱۳۹۹
		"{YYY}", enToFa.Replace(strconv.Itoa(c.GetYear())[1:]),
		"{YY}", enToFa.Replace(strconv.Itoa(c.GetYear())[2:]),
		"{Y}", enToFa.Replace(strconv.Itoa(c.GetYear())[3:]),
		// Month
		"{MM}", enToFa.Replace(fmt.Sprintf("%02d", c.GetMonth())),
		"{M}", enToFa.Replace(fmt.Sprintf("%d", c.GetMonth())),
		// string
		"{MO}", toPersian(c.GetMonth()),
		// Day
		"{DD}", enToFa.Replace(fmt.Sprintf("%02d", c.GetDay())),
		"{D}", enToFa.Replace(fmt.Sprintf("%d", c.GetDay())),
		"{WD}", wdToFa.Replace(fmt.Sprintf("%d", c.GetWeekday())),
		// Hour
		"{HH}", enToFa.Replace(fmt.Sprintf("%02d", c.GetHour())),
		"{H}", enToFa.Replace(fmt.Sprintf("%d", c.GetHour())),
		// Minute
		"{TT}", enToFa.Replace(fmt.Sprintf("%02d", c.GetMinute())),
		"{T}", enToFa.Replace(fmt.Sprintf("%d", c.GetMinute())),
		// Second
		"{SS}", enToFa.Replace(fmt.Sprintf("%02d", c.GetSecond())),
		"{S}", enToFa.Replace(fmt.Sprintf("%d", c.GetSecond())),
	)
	return r.Replace(format)
}
func toPersian(month int) string {
	replacer := strings.NewReplacer(
		"01", "فروردین",
		"02", "اردیبهشت",
		"03", "خرداد",
		"04", "تیر",
		"05", "مرداد",
		"06", "شهریور",
		"07", "مهر",
		"08", "آبان",
		"09", "آذر",
		"10", "دی",
		"11", "بهمن",
		"12", "اسفند")
	return replacer.Replace(fmt.Sprintf("%02d", month))
}

func toEng(month int) string {
	replacer := strings.NewReplacer(
		"01", "Farvardin",
		"02", "Ordibehesht",
		"03", "Khordad",
		"04", "Tir",
		"05", "Mordad",
		"06", "Shahrivar",
		"07", "Mehr",
		"08", "Aban",
		"09", "Azar",
		"10", "Dey",
		"11", "Bahman",
		"12", "Esfand",
		)
	return replacer.Replace(strconv.Itoa(month))
}

var enToFa = strings.NewReplacer(
	"0", "۰",
	"1", "۱",
	"2", "۲",
	"3", "۳",
	"4", "۴",
	"5", "۵",
	"6", "۶",
	"7", "۷",
	"8", "۸",
	"9", "۹",
)

var wdToFa = strings.NewReplacer(
	"Saturday", "6",
	"0", "یکشنبه",
	"1", "دوشنبه",
	"2", "سه شنبه",
	"3", "چهارشنبه",
	"4", "پنج شنبه",
	"5", "جمعه",
)

var wdToEn = strings.NewReplacer(
	"6", "Saturday",
	"0", "Sunday",
	"1", "Monday",
	"2", "Tuesday",
	"3", "Wednesday",
	"4", "Thursday",
	"5", "Friday",
)

//  {days, leap_days, days_before_start}
//var pMonthCount = [12][3]int{
//	{31, 31, 0},   // Farvardin
//	{31, 31, 31},  // Ordibehesht
//	{31, 31, 62},  // Khordad
//	{31, 31, 93},  // Tir
//	{31, 31, 124}, // Mordad
//	{31, 31, 155}, // Shahrivar
//	{30, 30, 186}, // Mehr
//	{30, 30, 216}, // Aban
//	{30, 30, 246}, // Azar
//	{30, 30, 276}, // Dey
//	{30, 30, 306}, // Bahman
//	{29, 30, 336}, // Esfand
//}


