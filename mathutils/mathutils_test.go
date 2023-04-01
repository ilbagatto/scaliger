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

func TestShortPolynome(t *testing.T) {
	got := Polynome(10, 1, 2, 3)
	exp := 321.0
	if !AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestLongPolynome(t *testing.T) {
	got := Polynome(
		-0.127296372347707,
		0.409092804222329,
		-0.0226937890431606,
		-7.51461205719781e-06,
		0.0096926375195824,
		-0.00024909726935408,
		-0.00121043431762618,
		-0.000189319742473274,
		3.4518734094999e-05,
		0.000135117572925228,
		2.80707121362421e-05,
		1.18779351871836e-05)
	exp := 0.411961500152426
	if !AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}
