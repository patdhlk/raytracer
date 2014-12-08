package objects

import (
	"testing"
)

func TestIntersection(t *testing.T) {
	sphere := NewSphere(NewVector(0.0, 0.0, 0.0), 1.0)
	ray := NewRay(NewVector(3.0, 0.0, 0.0), NewVector(-1.0, 0.0, 0.0))
	pos := sphere.Intersection(ray)
	if pos.X() != 1.0 || pos.Y() != 0.0 || pos.Z() != 0.0 {
		t.Errorf("TestIntersection %v %v %v)", pos.X(), pos.Y(), pos.Z())
	}
}
