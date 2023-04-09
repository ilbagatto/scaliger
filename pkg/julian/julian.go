package julian

// The main purpose is to convert between civil dates and Julian dates.
// Julian date (JD) is the number of days elapsed since mean UT noon of
// January 1st 4713 BC. This system of time measurement is widely adopted by
// the astronomers.

import (
	"math"
)

const SEC_PER_DAY = 24 * 60 * 60
const DAYS_PER_SEC = 1.0 / SEC_PER_DAY // How many days in a second?
const DAYS_PER_CENT = 36525
const J2000 = 2451545.0 // Julian day for 2000 Jan. 1.5
const J1900 = 2415020.0 // Julian day for 1900 Jan. 0.5

func isGregorian(date CivilDate) bool {
	if date.Year > 1582 {
		return true
	}
	if date.Year < 1582 {
		return false
	}
	if date.Month > 10 {
		return true
	}
	if date.Month < 10 {
		return false
	}
	if date.Day > 10 {
		return true
	}
	return false
}

// Converts calendar date into Julian days.
func CivilToJulian(date CivilDate) float64 {
	var y, m float64
	if date.Month > 2 {
		y = float64(date.Year)
		m = float64(date.Month)
	} else {
		y = float64(date.Year) - 1
		m = float64(date.Month) + 12
	}

	var t float64
	if date.Year < 0 {
		t = 0.75
	}

	var a, b float64
	if isGregorian(date) {
		a = math.Trunc(y / 100)
		b = 2 - a + math.Trunc(a/4)
	}

	return b + math.Trunc(365.25*y-t) + math.Trunc(30.6001*(m+1)) + date.Day + 1720994.5
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

	return CivilDate{Year: int(ye), Month: int(mo), Day: da}
}

// Given number of Julian days, calculates JD at Greenwich midnight.
func JulianMidnight(jd float64) float64 {
	return math.Floor(jd-0.5) + 0.5
}

// Julian Day corresponding to January 0.0 of a given year
func JulianDateZero(year int) float64 {
	y := float64(year - 1)
	a := math.Trunc(y / 100)
	return math.Trunc(365.25*y) - a + math.Trunc(a/4) + 1721424.5
}

// Converts fractional part of a Julian Date to decimal hours.
func ExtractUTC(jd float64) float64 {
	// _, f := math.Modf(jd - 0.5)
	// return math.Abs(f) * 24.0
	return (jd - JulianMidnight(jd)) * 24
}
