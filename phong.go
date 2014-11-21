package raytracer

import "math"

func calcIa(in, ka float64) float64 {
	return in * ka
}

func calcId(in, kd float64, L, N vector) float64 {
	return in * kd * Multiply(L, N)
}

func calcIs(in, ks float64, b, V vector, n float64) float64 {
	return in * ks * math.Pow(-Multiply(b, V), n)
}

func CalcBrightness(in float64, b, L, N, V vector, m material, n float64) float64 {
	return calcIa(in, m.getKa()) + calcId(in, m.getKd(), L, N) + calcIs(in, m.getKs(), b, V, n)
}
