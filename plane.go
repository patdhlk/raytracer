package raytracer

import (
	"image/color"
)

//Ebene
type Plane struct {
	Point  Vector
	Normal Vector
	Color  color.RGBA
}

func NewPlane(point Vector, normal Vector, color color.RGBA) *Plane {
	p := Plane{point, normal, color}
	return &p
}

func (p *Plane) findIntersection(r Ray) (int32, float64, float64) {
	return 1, p.Point.AddVector(*r.Origin.Negative()).DotProduct(p.Normal) / r.Direction.DotProduct(p.Normal), 0.0
}

func (p *Plane) getNormalAt(point Vector) *Vector {
	return NewVectorByVector(p.Normal)
}
