package sidereal

// Converts civil time into the sidereal time and vice-versa.
//
// Some functions expect [year], [month] and [day] as input arguments.
// Thoey always represent a date of proleptic Gregorian
// calendar. The [year] is astronomical, negative for BC dates.
// Zero year corresponds to -1 BC.

import (
	"github.com/skrushinsky/scaliger/mathutils"
)

// Converts Julian date to Mean Greenwich Sidereal Time
func JulianToSidereal(jd, lng float64) float64 {
	d := jd - 2451545.0
	t := d / 36525
	t2 := t * t
	t3 := t2 * t
	gst := 280.46061837 + 360.98564736629*d + 0.000387933*t2 - t3/38710000
	return mathutils.ReduceHours((gst + lng) / 15)
}
