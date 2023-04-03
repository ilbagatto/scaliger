package nutequ

// Given [jd], number of Julian days, calculate **obliquity of ecliptic**
// in degrees.
func MeanObliquity(jd float64) float64 {
	t := (jd - 2415020) / 36525
	c := (((-0.00181*t)+0.0059)*t + 46.845) * t
	return 23.45229444 - (c / 3600)
}
