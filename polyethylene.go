package main

type Polyethylene struct {
}

func (p Polyethylene) getKa() float64 {
	return 0.7
}

func (p Polyethylene) getKd() float64 {
	return 0.8
}

func (p Polyethylene) getKs() float64 {
	return 0.9
}

func (p Polyethylene) getN() int {
	return 30
}
