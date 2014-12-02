package main

import "testing"

func TestCalcReflecion(t *testing.T) {
	b := NewVector(1.0, 2.0, 3.0)
	N := NewVector(4.0, 5.0, 6.0)

	result := CalcReflecion(b, N)

	if result.x != -255.0 || result.y != -318.0 || result.z != -381.0 {
		t.Errorf("TestCalcReflecion %v %v %v %v %v %v", -255.0, -318.0, -381.0, result.x, result.y, result.z)
	}
}
