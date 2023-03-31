package sidereal

// Converts a given civil time into the local sidereal time and vice-versa.

import (
	"github.com/skrushinsky/scaliger/julian"
	"github.com/skrushinsky/scaliger/mathutils"
)

func CivilToSidereal(date julian.CivilDate) float64 {

	jd := julian.CivilToJulian(date)
	jdm := julian.JulianMidnight(jd)
	jd0 := julian.JulianDateZero(date.Year)
	days := jdm - jd0

	t := (jd0 - 2415020) / 36525
	r := 6.6460656 + 2400.051262*t + 0.00002581*t*t
	b := 24 - r + float64(24*(date.Year-1900))
	t0 := 0.0657098*days - b
	ut := mathutils.Frac(date.Day) * 24
	gst := t0 + 1.002738*ut
	// TODO: use range function
	gst = mathutils.ReduceHours(gst)
	return gst
}
