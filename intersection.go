package raytracer

import "math"

func CalcIntersectionLineBall(l line, b ball) (int32, float64, float64) {
	var d vector
	var e float64
	var f float64
	var g float64
	var resCount int32 = 0
	var t1 float64 = 0
	var t2 float64 = 0

	d = Subtract(l.a, b.k)

	f = float64(Multiply(d, l.b))
	e = math.Pow(f, 2.0) + math.Pow(float64(b.R), 2.0) - math.Pow(float64(Multiply(d, d)), 2.0)

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

func CalcIntersectionPlaneLine(p plane, l line) int32 {
	return Multiply(Subtract(p.e, l.a), p.n) / Multiply(l.b, p.n)
}
