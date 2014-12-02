package main

import (
	"image/color"
	"math"
	"testing"
)

func TestNewSphere(t *testing.T) {
	x, y, z := 4.0, 2.0, 8.0
	vec := NewVector(x, y, z)

	r := 4.0

	s := NewSphere(vec, r, color.RGBA{255, 255, 255, 0})

	if s.location.x != x || s.location.y != y || s.location.z != z || s.radius != r {
		t.Errorf("TestNewSphere %v %v %v %v %v %v %v %v", x, y, z, r, s.location.x, s.location.y, s.location.z, s.radius)
	}
}

func TestSphereFindIntersection(t *testing.T) {
	c := color.RGBA{255, 255, 255, 0}
	s := NewSphere(NewVector(0.0, 0.0, 0.0), 4.0, c)
	r := NewRay(NewVector(0.0, 0.0, 0.0), NewVector(0.0, 0.0, 0.0))
	distance, color, res := s.FindIntersection(r)

	if res != true || distance != 4.0 {
		t.Errorf("TestSphereFindIntersection1 %v %v %v", res, distance, color)
	}

	reflection, err := s.GetReflectionRay(r, 1)

	if err != nil || reflection.origin != NewVector(0, 0, 0) || reflection.direction != NewVector(0, 0, 0) {
		t.Errorf("TestSphereFindIntersection ReflectionRay 1 %v %v", reflection, err)
	}

	s = NewSphere(NewVector(0.0, 0.0, 0.0), 4.0, c)
	r = NewRay(NewVector(2.0, 0.0, 0.0), NewVector(0.0, 0.0, 0.0))
	distance, color, res = s.FindIntersection(r)

	if res != true || distance != math.Sqrt(12.0) {
		t.Errorf("TestSphereFindIntersection2 %v %v %v", res, distance, color)
	}

	reflection, err = s.GetReflectionRay(r, 1)

	if err != nil || reflection.origin != NewVector(2, 0, 0) || reflection.direction != NewVector(0, 0, 0) {
		t.Errorf("TestSphereFindIntersection ReflectionRay 2 %v %v", reflection, err)
	}

	s = NewSphere(NewVector(0.0, 0.0, 0.0), 4.0, c)
	r = NewRay(NewVector(3.0, 0.0, 0.0), NewVector(0.0, 0.0, 0.0))
	distance, color, res = s.FindIntersection(r)

	if res != true || distance != math.Sqrt(7.0) {
		t.Errorf("TestSphereFindIntersection3 %v %v %v", res, distance, color)
	}

	reflection, err = s.GetReflectionRay(r, 1)

	if err != nil || reflection.origin != NewVector(3, 0, 0) || reflection.direction != NewVector(0, 0, 0) {
		t.Errorf("TestSphereFindIntersection ReflectionRay 3 %v %v", reflection, err)
	}

	s = NewSphere(NewVector(0.0, 0.0, 0.0), 4.0, c)
	r = NewRay(NewVector(5.0, 0.0, 0.0), NewVector(0.0, 1.0, 0.0))
	distance, color, res = s.FindIntersection(r)

	if res != false || distance != math.Sqrt(0.0) {
		t.Errorf("TestSphereFindIntersection4 %v %v %v", res, distance, color)
	}

	reflection, err = s.GetReflectionRay(r, 1)

	if err == nil || reflection.origin != NewVector(0, 0, 0) || reflection.direction != NewVector(0, 0, 0) {
		t.Errorf("TestSphereFindIntersection ReflectionRay 4 %v %v", reflection, err)
	}
}

func TestSphereGetNormalAt(t *testing.T) {
	s := NewSphere(NewVector(0.0, 0.0, 0.0), 4.0, color.RGBA{255, 255, 255, 0})

	v := s.GetNormalAt(NewVector(0, 2, 0))

	if v != NewVector(0, 1, 0) {
		t.Errorf("TestSphereGetNormalAt1 %v", v)
	}

	s = NewSphere(NewVector(2.0, 2.0, 2.0), 4.0, color.RGBA{255, 255, 255, 0})

	v = s.GetNormalAt(NewVector(2, 6, 2))

	if v != NewVector(0, 1, 0) {
		t.Errorf("TestSphereGetNormalAt2 %v", v)
	}
}

func TestSphereGetReflectionRay(t *testing.T) {
	//s = NewSphere(*NewVector(0.0, 0.0, 0.0), 4.0, 0.0)

	//
}
