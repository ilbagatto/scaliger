package nutequ

import (
	"testing"

	"github.com/skrushinsky/scaliger/mathutils"
)

type _OblTestCase struct {
	jd  float64
	eps float64
}

var epsCases = [...]_OblTestCase{
	{
		jd:  2444140.5, // 1979-09-24.0
		eps: 23.441916666666668,
	},
	{
		jd:  2451544.5, // 2000-01-01.0
		eps: 23.43927777777778,
	},
}

func TestMeamObliquity(t *testing.T) {
	for _, test := range epsCases {
		eps := MeanObliquity(test.jd)
		if !mathutils.AlmostEqual(eps, test.eps, 1e-4) {
			t.Errorf("Expected: %f, got: %f", test.eps, eps)
		}
	}
}
