package sidereal

import (
	"testing"

	"github.com/skrushinsky/scaliger/julian"
	"github.com/skrushinsky/scaliger/mathutils"
)

func TestCivilToSidereal(t *testing.T) {
	got := CivilToSidereal(julian.CivilDate{Year: 2010, Month: 2, Day: 7.97917})
	exp := 8.698091
	if !mathutils.AlmostEqual(got, exp, 1e-4) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}
