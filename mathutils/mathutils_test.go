package mathutils

import "testing"

func TestPositiveFix(t *testing.T) {
	got := Fix(1.5)
	if !AlmostEqual(got, 1.0, 1e-6) {
		t.Errorf("Expected 1.0, got: %f", got)
	}
}

func TestNegativeFix(t *testing.T) {
	got := Fix(-1.5)
	if !AlmostEqual(got, -1.0, 1e-6) {
		t.Errorf("Expected -1.0, got: %f", got)
	}
}
