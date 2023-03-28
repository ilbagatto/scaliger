package scaliger

import (
	"testing"
)

func equalDates(a, b CivilDate) bool {
	return a.year == b.year && a.month == b.month && AlmostEqual(a.date, b.date, 1e-6)
}

func TestCivilToJulianAfterGregorian(t *testing.T) {
	date := CivilDate{year: 2010, month: 1, date: 1.0}
	exp := 2455197.5
	got := CivilToJulian(date)
	if !AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestCivilToJulianBeforeGregorian(t *testing.T) {
	date := CivilDate{year: 837, month: 4, date: 10.3}
	exp := 2026871.8
	got := CivilToJulian(date)
	if !AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestCivilToJulianBC(t *testing.T) {
	date := CivilDate{year: -1000, month: 7, date: 12.5}
	exp := 1356001.0
	got := CivilToJulian(date)
	if !AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestJulianToCivilAfterGregorian(t *testing.T) {
	jd := 2455197.5
	exp := CivilDate{year: 2010, month: 1, date: 1.0}
	got := JulianToCivil(jd)
	if !equalDates(got, exp) {
		t.Errorf("Expected: %d-%d-%f, got: %d-%d-%f", exp.year, exp.month, exp.date, got.year, got.month, got.date)
	}
}

func TestJulianToCivilBeforeGregorian(t *testing.T) {
	jd := 2026871.8
	exp := CivilDate{year: 837, month: 4, date: 10.3}
	got := JulianToCivil(jd)
	if !equalDates(got, exp) {
		t.Errorf("Expected: %d-%d-%f, got: %d-%d-%f", exp.year, exp.month, exp.date, got.year, got.month, got.date)
	}
}

func TestJulianToCivilBC(t *testing.T) {
	jd := 1356001.0
	exp := CivilDate{year: -1000, month: 7, date: 12.5}
	got := JulianToCivil(jd)
	if !equalDates(got, exp) {
		t.Errorf("Expected: %d-%d-%f, got: %d-%d-%f", exp.year, exp.month, exp.date, got.year, got.month, got.date)
	}
}
