package sidereal

import (
	"testing"

	"github.com/skrushinsky/scaliger/mathutils"
)

type SidTestCase struct {
	jd  float64
	lst float64
}

var cases = [...]SidTestCase{
	{jd: 2445943.851053, lst: 7.072111}, // 1984-08-31.4
	{jd: 2415703.498611, lst: 3.525306}, // 1901-11-15.0
	{jd: 2415702.501389, lst: 3.526444}, // 1901-11-14.0
	{jd: 2444352.108931, lst: 4.668119}, // 1980-04-22.6
}

func TestJulianToSidereal(t *testing.T) {

	for _, test := range cases {
		lst := JulianToSidereal(test.jd, 0)
		if !mathutils.AlmostEqual(lst, test.lst, 1e-4) {
			t.Errorf("Expected: %f, got: %f", test.lst, lst)
		}
	}
}
