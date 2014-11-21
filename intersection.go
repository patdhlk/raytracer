package raytracer

import "math"

func CalcIntersection(a, b, k vector, r int32) (int32, float64, float64) {
	var d vector
	var e float64
	var f float64
	var g float64
	var resCount int32 = 0
	var t1 float64 = 0
	var t2 float64 = 0
	d = Subtract(a, k)

	f = float64(Multiply(d, b))
	e = math.Pow(f, 2.0) + math.Pow(float64(r), 2.0) - math.Pow(float64(Multiply(d, d)), 2.0)

	if e >= 0 {
		if e == 0 {
			resCount = 1
			t1 = -f
		} else {
			g = math.Sqrt(e)
			resCount = 2
			t1 = -f + g
			t2 = -f - g
		}
	}
	return resCount, t1, t2
}
