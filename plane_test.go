package raytracer

import (
	"image/color"
	"testing"
)

func TestPlaneFindIntersection(t *testing.T) {
	p := NewPlane(*NewVector(0, 0, 0), *NewVector(0, 1, 0), color.RGBA{255, 255, 255, 0})

	t1 := p.findIntersection(*NewRay(*NewVector(0, 0, 0), *NewVector(0, -1, 0)))

	if t1 != 0.0 {
		t.Errorf("TestPlaneFindIntersection1 %v", t1)
	}

	t1 = p.findIntersection(*NewRay(*NewVector(0, 1, 0), *NewVector(0, -1, 0)))

	if t1 != 1.0 {
		t.Errorf("TestPlaneFindIntersection2 %v", t1)
	}

	t1 = p.findIntersection(*NewRay(*NewVector(0, -1, 0), *NewVector(0, -1, 0)))

	if t1 != -1.0 {
		t.Errorf("TestPlaneFindIntersection3 %v", t1)
	}
}
