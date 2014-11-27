package raytracer

import (
	"image/color"
	"math"
)

//Ball/Kugel
type Sphere struct {
	Center Vector
	Radius float64
	Color  color.RGBA
}

func NewSphere(center Vector, radius float64, color color.RGBA) *Sphere {
	//if (radius > 0.0)
	s := Sphere{center, radius, color}
	return &s
}

func (s *Sphere) findIntersection(r Ray) (int32, float64, float64) {
	var e, f, g, t1, t2 float64
	var resCount int32 = 0
	t1 = 0
	t2 = 0

	var d = *r.Origin.AddVector(*s.Center.Negative())

	f = d.DotProduct(r.Direction)
	e = math.Pow(f, 2.0) + math.Pow(s.Radius, 2.0) - math.Pow(d.DotProduct(d), 2.0)

	if e >= 0 {
		if e == 0 {
			resCount = 1
			t1 = -f
		} else {
			g = math.Sqrt(e)
			resCount = 2
			t1 = -f + g
			t2 = -f - g
		}
	}
	return resCount, t1, t2
}

func (s *Sphere) getNormalAt(point Vector) *Vector {
	return point.AddVector(*s.Center.Negative()).Normalize()
}
