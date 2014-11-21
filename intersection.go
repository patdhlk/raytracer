package raytracer

import "math"

func CalcIntersectionLineBall(l line, s sphere) (int32, float64, float64) {
	var d vector
	var e, f, g, t1, t2 float64
	var resCount int32 = 0
	t1 = 0
	t2 = 0

	d = Subtract(l.a, s.center)

	f = Multiply(d, l.b)
	e = math.Pow(f, 2.0) + math.Pow(s.radius, 2.0) - math.Pow(Multiply(d, d), 2.0)

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

func CalcIntersectionPlaneLine(p plane, l line) float64 {
	return Multiply(Subtract(p.e, l.a), p.n) / Multiply(l.b, p.n)
}
