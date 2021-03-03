package jalali


// A month() specifies a month() of the year() starting from Farvardin = 1.
type Month int

// A Weekday specifies a dey() of the week starting from Shanbe = 0.
type Weekday int

// A AmPm specifies the 12-Hour marker.
type AmPm int

// List of months in Persian JalaliCal.
const (
	Farvardin Month = 1 + iota
	Ordibehesht
	Khordad
	Tir
	Mordad
	Shahrivar
	Mehr
	Aban
	Azar
	Dey
	Bahman
	Esfand
)

// List of Dari months in Persian JalaliCal.
const (
	Hamal Month = 1 + iota
	Sur
	Jauza
	Saratan
	Asad
	Sonboleh
	Mizan
	Aqrab
	Qos
	Jady
	Dolv
	Hut
)

// List of days in a week.
const (
	Shanbeh Weekday = iota
	Yekshanbeh
	Doshanbeh
	Seshanbeh
	Charshanbeh
	Panjshanbeh
	Jomeh
)

// List of 12-Hour markers.
const (
	Am AmPm = 0 + iota
	Pm
)

var amPm = [2]string{
	"قبل از ظهر",
	"بعد از ظهر",
}

var sAmPm = [2]string{
	"ق.ظ",
	"ب.ظ",
}

var months = [12]string{
	"فروردین",
	"اردیبهشت",
	"خرداد",
	"تیر",
	"مرداد",
	"شهریور",
	"مهر",
	"آبان",
	"آذر",
	"دی",
	"بهمن",
	"اسفند",
}

var dmonths = [12]string{
	"حمل",
	"ثور",
	"جوزا",
	"سرطان",
	"اسد",
	"سنبله",
	"میزان",
	"عقرب",
	"قوس",
	"جدی",
	"دلو",
	"حوت",
}

var days = [7]string{
	"شنبه",
	"یک‌شنبه",
	"دوشنبه",
	"سه‌شنبه",
	"چهارشنبه",
	"پنج‌شنبه",
	"جمعه",
}

var sdays = [7]string{
	"ش",
	"ی",
	"د",
	"س",
	"چ",
	"پ",
	"ج",
}
