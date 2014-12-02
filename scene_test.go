package main

import (
	"image/color"
	"testing"
)

func TestNewScene(t *testing.T) {
	e := Eye{NewVector(1, 2, 3)}
	l := NewLight(NewVector(2, 3, 4))
	o := make([]SceneObject, 0)
	o = append(o, NewSphere(NewVector(4, 9, 14), 9.0, color.RGBA{200, 100, 0, 255}))

	s := NewScene(e, l, o)

	if s.eye.Point != NewVector(1, 2, 3) || s.light.Position != NewVector(2, 3, 4) || len(s.objects) != 1 {
		t.Errorf("TestNewScene %v %v %v", s.eye, s.light, s.objects)
	}
}

func TestNewDefaultScene(t *testing.T) {
	s := NewDefaultScene()

	if &s.eye == nil || &s.light == nil || s.objects == nil || len(s.objects) <= 0 {
		t.Errorf("TestNewDefaultScene %v %v %v", s.eye, s.light, s.objects)
	}
}
