// Converts civil time into the sidereal time.
//
// Source: Peter Duffett-Smith, "Astronomy with your PC", 2-d edition
package sidereal

import (
	"math"

	"github.com/skrushinsky/scaliger/julian"
	"github.com/skrushinsky/scaliger/mathutils"
)

const SOLAR_TO_SIDEREAL = 1.002737909350795

type SiderealOptions struct {
	Lng  float64 // geographical longitude, degrees, negative westwards
	Eps  float64 // obliquity of the ecliptic, degrees
	Dpsi float64 // nutation in longitude, degrees
}

func meanGMST(jd float64) float64 {
	date := julian.JulianToCivil(jd)
	dj := julian.JulianMidnight(jd) - julian.J1900
	t := dj/julian.DAYS_PER_CENT - 1
	t2 := t * t
	t3 := t * t2
	r1 := 6.697374558 + (2400 * (t - (float64(date.Year-2000) / 100)))
	r0 := (5.13366e-2 * t) + (2.586222e-5 * t2) - (1.722e-9 * t3)
	t0 := mathutils.ReduceHours(r0 + r1)
	return julian.ExtractUTC(jd)*SOLAR_TO_SIDEREAL + t0
}

// Converts Julian date to Sidereal Time.
// If options contain initialized [Lng] field, then the result is  **Local Sidereal Time**.
// Otherwise, Greenwich Sidereal Time.
//
// If options contains initialized [Eps] and [Dpsi] fields, then the result is
// **apparent Sidereal Time**. Otherwise, **Mean Sidereal Time**.
func JulianToSidereal(jd float64, options SiderealOptions) float64 {
	dpsi := options.Dpsi * 3600                                       // degrees -> arcseconds
	delta := (dpsi * math.Cos(mathutils.Radians(options.Eps))) / 15.0 // correction in seconds of time
	lng := options.Lng / 15
	return mathutils.ReduceHours(meanGMST(jd) + delta/3600 + lng)
}
