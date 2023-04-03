package nutequ

import (
	"testing"

	"github.com/skrushinsky/scaliger/mathutils"
)

type _NutTestCase struct {
	jd   float64
	dpsi float64
	deps float64
}

var nutCases = [...]_NutTestCase{
	{
		jd:   2399215.5, // 1856 Sept. 23
		dpsi: -0.00127601021242336,
		deps: 0.00256293723137559,
	},
	{
		jd:   2451544.5, // 2000 Jan. 1
		dpsi: -0.00387728730373955,
		deps: -0.00159919822661103,
	},
	{
		jd:   2443825.69, // 1978 Nov 17
		dpsi: -9.195562346652888e-04,
		deps: -2.635113483663831e-03,
	},
}

func TestNutLon(t *testing.T) {
	for _, test := range nutCases {
		dpsi, _ := Nutation(test.jd)
		if !mathutils.AlmostEqual(dpsi, test.dpsi, 1e-4) {
			t.Errorf("Expected: %f, got: %f", test.dpsi, dpsi)
		}
	}
}

func TestNutEcl(t *testing.T) {
	for _, test := range nutCases {
		_, deps := Nutation(test.jd)
		if !mathutils.AlmostEqual(deps, test.deps, 1e-4) {
			t.Errorf("Expected: %f, got: %f", test.deps, deps)
		}
	}
}
