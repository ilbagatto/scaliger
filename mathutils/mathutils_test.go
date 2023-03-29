package mathutils

import "testing"

func TestAlmostEqualPositive(t *testing.T) {
	if !AlmostEqual(10.1234567, 10.1234568, 1e-6) {
		t.Errorf("Numbers should be considered equal")
	}
}

func TestAlmostEqualNegative(t *testing.T) {
	if AlmostEqual(10.123457, 10.123456, 1e-6) {
		t.Errorf("Numbers should be considered equal")
	}
}
