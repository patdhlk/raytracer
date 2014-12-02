package main

import "testing"

func TestNewRay(t *testing.T) {
	r := NewRay(NewVector(0, 1, 2), NewVector(3, 4, 0))

	if r.origin.x != 0.0 || r.origin.y != 1.0 || r.origin.z != 2.0 || r.direction.x != 0.6 || r.direction.y != 0.8 || r.direction.z != 0.0 {
		t.Errorf("TestNewScreen %v %v", r.origin, r.direction)
	}
}
