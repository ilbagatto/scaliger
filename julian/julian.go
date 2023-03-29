package julian

// The main purpose is to convert between civil dates and Julian dates.
// Julian date (JD) is the number of days elapsed since mean UT noon of
// January 1st 4713 BC. This system of time measurement is widely adopted by
// the astronomers.

import (
	"math"
)

// Civil calendar date, usually Gregorian.
// Time is represented by fractional part of a day.
// For example, 7h30m UT is `(7 + 30 / 60) / 24 = 0.3125`.
type CivilDate struct {
	year  int     // a year, astronomical, negative for BC dates
	month int     // a month number, 1-12
	date  float64 // date which fractional part represents hours
}

func isGregorian(date CivilDate) bool {
	if date.year > 1582 {
		return true
	}
	if date.year < 1582 {
		return false
	}
	if date.month > 10 {
		return true
	}
	if date.month < 10 {
		return false
	}
	if date.date > 10 {
		return true
	}
	return false
}

// Converts calendar date into Julian days.
func CivilToJulian(date CivilDate) float64 {
	var y, m float64
	if date.month > 2 {
		y = float64(date.year)
		m = float64(date.month)
	} else {
		y = float64(date.year) - 1
		m = float64(date.month) + 12
	}

	var t float64
	if date.year < 0 {
		t = 0.75
	}

	var a, b float64
	if isGregorian(date) {
		a = math.Trunc(y / 100)
		b = 2 - a + math.Trunc(a/4)
	}

	return b + math.Trunc(365.25*y-t) + math.Trunc(30.6001*(m+1)) + date.date + 1720994.5
}

// Converts [jd], number of Julian days into the calendar date.
func JulianToCivil(jd float64) CivilDate {
	i, f := math.Modf(jd + 0.5)

	var b float64
	if i > 2299160 {
		a := math.Trunc((i - 1867216.25) / 36524.25)
		b = i + 1 + a - math.Trunc(a/4)
	} else {
		b = i
	}
	c := b + 1524
	d := math.Trunc((c - 122.1) / 365.25)
	e := math.Trunc(365.25 * d)
	g := math.Trunc((c - e) / 30.6001)

	da := c - e + f - math.Trunc(30.6001*g)
	var mo float64
	if g < 13.5 {
		mo = g - 1
	} else {
		mo = g - 13
	}
	var ye float64
	if mo > 2.5 {
		ye = d - 4716
	} else {
		ye = d - 4715
	}

	return CivilDate{int(ye), int(mo), da}
}

// Given number of Julian days, calculates JD at Greenwich midnight.
func JulianMidnight(jd float64) float64 {
	return math.Floor(jd-0.5) + 0.5
}
