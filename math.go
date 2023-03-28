package scaliger

// Core mathematical functions for astronomical calculations.

import "math"

// Returns the integer part of an [x]
// [Fix(1.5) = 1]
// [Fix(-1.5) = -1]
func Fix(x float64) float64 {
	if x < 0 {
		return math.Ceil(x)
	}
	return math.Floor(x)
}

// Compares two floats with a given precision.
func AlmostEqual(a, b, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}
