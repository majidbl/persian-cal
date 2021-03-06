package jalali

import "testing"

func TestDate(t *testing.T) {
	test := Date(1399,12,16,12,12,12)
	expect := Calendar{
		year:    1399,
		month:   12,
		day:     16,
		hour:    12,
		min:     12,
		sec:     12,
		weekDay: 0,
		loc:     nil,
	}
	if test != expect{
		t.Fatalf("expected %v founded %v", expect, test)
	}
}

func TestToJalali(t *testing.T) {
	test, err := ToJalali(2021,3,6)
	if err != nil{
		t.Error(err)
	}
	expected := Date(1399,12,16,0,0,0)
	if test != expected{
		t.Fatalf("expected %v founded %v", expected, test)
	}
}

func TestToGregorian(t *testing.T) {
	date := Date(1399,12,16,0,0,0)
	georgianDate, err := date.ToGregorian()
	if err != nil {
		t.Error(err)
	}
	expected := Date(2021,3,6,0,0,0)
	if georgianDate != expected{
		t.Fatalf("expected %v founded %v", expected, georgianDate)
	}
}