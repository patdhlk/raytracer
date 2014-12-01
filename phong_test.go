package main

import "testing"

type PhongTestMaterial struct {
}

func (p PhongTestMaterial) getKa() float64 {
	return 0.7
}

func (p PhongTestMaterial) getKd() float64 {
	return 0.8
}

func (p PhongTestMaterial) getKs() float64 {
	return 0.9
}

func (p PhongTestMaterial) getN() int {
	return 2
}

func TestCalcPhong(t *testing.T) {
	p := PhongTestMaterial{}

	//func calcIa(in, ka float64) float64 {
	ia := calcIa(2.0, 0.7)

	if ia != 1.4 {
		t.Errorf("TestCalcPhong ia %v %v", 1.4, ia)
	}

	//func calcId(in, kd float64, L, N Vector) float64 {
	id := calcId(2.0, 0.8, NewVector(0.6, 0.8, 0.0), NewVector(0.6, 0.8, 0.0))

	if id != 1.6 {
		t.Errorf("TestCalcPhong id %v %v", 1.6, id)
	}

	//func calcIs(in, ks float64, b, V Vector, n float64) float64 {
	is := calcIs(2.0, 0.9, NewVector(0.6, 0.8, 0.0), NewVector(0.8, 0.6, 0.0), 0)

	if is != 1.8 {
		t.Errorf("TestCalcPhong id 1 %v %v", 1.8, is)
	}

	is = calcIs(2.0, 0.9, NewVector(0.6, 0.8, 0.0), NewVector(0.8, 0.6, 0.0), 2)

	if is != 1.65888 {
		t.Errorf("TestCalcPhong id 2 %v %v", 1.65888, is)
	}

	phong := CalcBrightness(2.0, NewVector(0.6, 0.8, 0.0), NewVector(0.6, 0.8, 0.0), NewVector(0.6, 0.8, 0.0), NewVector(0.8, 0.6, 0.0), p)

	phongCalc := ia + id + is

	if phong != phongCalc {
		t.Errorf("TestCalcPhong %v %v", phongCalc, phong)
	}
}
