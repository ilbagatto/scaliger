// Core mathematical functions for astronomical calculations.
package mathutils

import "math"

const _DEG_RAD = math.Pi / 180
const _RAD_DEG = 180 / math.Pi

// Compares two floats with a given precision.
func AlmostEqual(a, b, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}

// Fractional part of a number.
//
// It uses the standard [math.Modf] function.
// The result always keeps sign of the argument.
//
//	Frac(-5.5) = -0.5
func Frac(x float64) float64 {
	_, f := math.Modf(x)
	return f
}

// Reduces hours to range 0 >= x < 24
func ReduceHours(hours float64) float64 {
	x := math.Mod(hours, 24)
	if x < 0 {
		return x + 24
	}
	return x
}

// Calculates polynome:
//
//	a1 + a2*t + a3*t*t + a4*t*t*t...
//
// t is a number of Julian centuries elapsed since 1900, Jan 0.5
// terms is a list of terms
func Polynome(t float64, terms ...float64) float64 {
	res := 0.0
	for i, k := range terms {
		p := math.Pow(t, float64(i))
		res += k * p
	}
	return res
}

// Converts arc-degrees to radians
func Radians(deg float64) float64 {
	return deg * _DEG_RAD
}

// Converts radians to arc-degrees
func Degrees(rad float64) float64 {
	return rad * _RAD_DEG
}

// Used with polinomial function for better accuracy.
func Frac360(x float64) float64 {
	f := Frac(x) * 360
	if f < 360 {
		f += 360
	}
	return f
}
