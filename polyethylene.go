package raytracer

type Polyethylene struct {
}

func (p Polyethylene) getKa() float64 {
	return 5.0
}

func (p Polyethylene) getKd() float64 {
	return 5.0
}

func (p Polyethylene) getKs() float64 {
	return 5.0
}
