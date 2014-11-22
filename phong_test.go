package raytracer

import "testing"

func TestCalcPhong(t *testing.T) {
	ia := calcIa(2.0, 3.5)

	//id := calcId()

	if ia != 7.0 {
		t.Errorf("TestCalcPhong %v %v", 7.0, calcId)
	}
}
