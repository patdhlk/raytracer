package main

import (
	"image/color"
	"testing"
)

func TestNewPlane(t *testing.T) {
	p := NewPlane(NewVector(1, 2, 3), NewVector(0, 1, 0), color.RGBA{255, 255, 255, 0})

	v := p.location
	if v.x != 1 || v.y != 2 || v.z != 3 {
		t.Errorf("TestNewPlane p.Point %v %v %v %v %v %v", 1, 2, 3, v.x, v.y, v.z)
	}

	v = p.normalDirection
	if v.x != 0 || v.y != 1 || v.z != 0 {
		t.Errorf("TestNewPlane p.Normal %v %v %v %v %v %v", 0, 1, 0, v.x, v.y, v.z)
	}
}

func TestPlaneFindIntersection(t *testing.T) {
	p := NewPlane(NewVector(0, 0, 0), NewVector(0, 1, 0), color.RGBA{255, 255, 255, 0})

	distance, color, res := p.FindIntersection(NewRay(NewVector(0, 0, 0), NewVector(0, -1, 0)))

	if res != true || distance != 0.0 {
		t.Errorf("TestPlaneFindIntersection1 %v %v %v", res, distance, color)
	}

	distance, color, res = p.FindIntersection(NewRay(NewVector(0, 1, 0), NewVector(0, -1, 0)))

	if res != false || distance != 0.0 {
		t.Errorf("TestPlaneFindIntersection2 %v %v %v", res, distance, color)
	}

	distance, color, res = p.FindIntersection(NewRay(NewVector(0, -1, 0), NewVector(0, -1, 0)))

	if res != true || distance != 1.0 {
		t.Errorf("TestPlaneFindIntersection3 %v %v %v", res, distance, color)
	}
}
