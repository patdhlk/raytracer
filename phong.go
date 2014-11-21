package raytracer

import "math"

func calcIa(in, ka float64) float64 {
	return in * ka
}

func calcId(in, kd float64, L, N Vector) float64 {
	return in * kd * L.DotProduct(N)
}

func calcIs(in, ks float64, b, V Vector, n float64) float64 {
	return in * ks * math.Pow(-(b.DotProduct(V)), n)
}

func CalcBrightness(in float64, b, L, N, V Vector, m material, n float64) float64 {
	return calcIa(in, m.getKa()) + calcId(in, m.getKd(), L, N) + calcIs(in, m.getKs(), b, V, n)
}
