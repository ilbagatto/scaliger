package mathutils

// Core mathematical functions for astronomical calculations.

import "math"

// Compares two floats with a given precision.
func AlmostEqual(a, b, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}
