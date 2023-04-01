package mathutils

// Core mathematical functions for astronomical calculations.

import "math"

// Compares two floats with a given precision.
func AlmostEqual(a, b, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}

// Fractional part of a number.
//
// It uses the standard [math.Modf] function.
// The result always keeps sign of the argument,
// e.g.: `Frac(-5.5) = -0.5`
func Frac(x float64) float64 {
	_, f := math.Modf(x)
	return f
}

// Reduces hours to range `0 >= x < 24`
func ReduceHours(hours float64) float64 {
	x := math.Mod(hours, 24)
	if x < 0 {
		return x + 24
	}
	return x
}

// Calculates polynome: `a1 + a2*t + a3*t*t + a4*t*t*t...`
// [t] s a number of Julian centuries elapsed since 1900, Jan 0.5
// [terms] is a list of terms
func Polynome(t float64, terms ...float64) float64 {
	res := 0.0
	for i, k := range terms {
		p := math.Pow(t, float64(i))
		res += k * p
	}
	return res
}
