package raytracer

import "testing"

func TestMaterialInterface(t *testing.T) {
	p := Polyethylene{}

	res := p.getKa()
	if res != 5 {
		t.Errorf("TestMaterialInterface ka %v %v", 5, res)
	}
	res = p.getKd()
	if res != 5 {
		t.Errorf("TestMaterialInterface kd %v %v", 5, res)
	}
	res = p.getKs()
	if res != 5 {
		t.Errorf("TestMaterialInterface ks %v %v", 5, res)
	}
}
