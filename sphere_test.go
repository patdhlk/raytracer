package raytracer

import "testing"

func TestNewSphere(t *testing.T) {
	x, y, z := 4.0, 2.0, 8.0
	vec := NewVector(x, y, z)

	r := 4.0

	s := NewSphere(*vec, r, 0.0)

	if s.Center.x != x || s.Center.y != y || s.Center.z != z || s.Radius != r {
		t.Errorf("TestNewSphere %v %v %v %v %v %v %v %v", x, y, z, r, s.Center.x, s.Center.y, s.Center.z, s.Radius)
	}
}

func TestFindIntersection(t *testing.T) {
	var s *Sphere
	var r *Ray
	var res int32
	var t1 float64
	var t2 float64

	s = NewSphere(*NewVector(0.0, 0.0, 0.0), 4.0, 0.0)
	r = NewRay(*NewVector(0.0, 0.0, 0.0), *NewVector(0.0, 0.0, 0.0))
	res, t1, t2 = s.findIntersection(*r)

	if res != 2 || t1 != 4.0 || t2 != -4.0 {
		t.Errorf("TestfindIntersection1 %v %v %v", res, t1, t2)
	}

	s = NewSphere(*NewVector(0.0, 0.0, 0.0), 4.0, 0.0)
	r = NewRay(*NewVector(2.0, 0.0, 0.0), *NewVector(0.0, 0.0, 0.0))
	res, t1, t2 = s.findIntersection(*r)

	if res != 1 || t1 != 0.0 || t2 != 0.0 {
		t.Errorf("TestfindIntersection2 %v %v %v", res, t1, t2)
	}

	s = NewSphere(*NewVector(0.0, 0.0, 0.0), 4.0, 0.0)
	r = NewRay(*NewVector(3.0, 0.0, 0.0), *NewVector(0.0, 0.0, 0.0))
	res, t1, t2 = s.findIntersection(*r)

	if res != 0 || t1 != 0.0 || t2 != 0.0 {
		t.Errorf("TestfindIntersection3 %v %v %v", res, t1, t2)
	}
}

func TestgetNormalAt(t *testing.T) {
	//s = NewSphere(*NewVector(0.0, 0.0, 0.0), 4.0, 0.0)

	//
}
