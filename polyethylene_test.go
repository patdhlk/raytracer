package main

import "testing"

func TestMaterialInterface(t *testing.T) {
	p := Polyethylene{}

	res := p.getKa()
	if res != 0.7 {
		t.Errorf("TestMaterialInterface ka %v %v", 0.7, res)
	}
	res = p.getKd()
	if res != 0.8 {
		t.Errorf("TestMaterialInterface kd %v %v", 0.8, res)
	}
	res = p.getKs()
	if res != 0.9 {
		t.Errorf("TestMaterialInterface ks %v %v", 0.9, res)
	}
	n := p.getN()
	if n != 30 {
		t.Errorf("TestMaterialInterface n %v %v", 30, n)
	}
}
