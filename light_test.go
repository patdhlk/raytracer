package main

import "testing"

func TestLight(t *testing.T) {
	l := NewLight(NewVector(1, 2, 3))

	v := l.Position
	if v != NewVector(1, 2, 3) {
		t.Errorf("TestLight %v %v", v, NewVector(1, 2, 3))
	}
}
