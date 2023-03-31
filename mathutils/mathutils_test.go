package mathutils

import "testing"

func TestAlmostEqualPositive(t *testing.T) {
	if !AlmostEqual(10.1234567, 10.1234568, 1e-6) {
		t.Errorf("Numbers should be considered equal")
	}
}

func TestAlmostEqualNegative(t *testing.T) {
	if AlmostEqual(10.123457, 10.123456, 1e-6) {
		t.Errorf("Numbers should not be considered equal")
	}
}

func TestPositiveFrac(t *testing.T) {
	if !AlmostEqual(Frac(5.5), 0.5, 1e-6) {
		t.Errorf("frac(5.5) should be 0.5")
	}
}

func TestNegativeFrac(t *testing.T) {
	if !AlmostEqual(Frac(-5.5), -0.5, 1e-6) {
		t.Errorf("frac(-5.5) should be -0.5")
	}
}

func TestReduceHoursPositive(t *testing.T) {
	if !AlmostEqual(ReduceHours(49.5), 1.5, 1e-6) {
		t.Errorf("49.5 should be reduced to 1.5")
	}
}

func TestReduceHoursNegative(t *testing.T) {
	if !AlmostEqual(ReduceHours(-0.5), 23.5, 1e-6) {
		t.Errorf("-0.5 should be reduced to 23.5")
	}
}
