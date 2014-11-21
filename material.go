package raytracer

type material interface {
	getKa() float64
	getKd() float64
	getKs() float64
}
