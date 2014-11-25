package raytracer

import (
	"image/color"
)

type Plane struct {
	Point  Vector
	Normal Vector
	Color  color.RGBA
}

func NewPlane(point Vector, normal Vector, color color.RGBA) *Plane {
	p := Plane{point, normal, color}
	return &p
}

func (p *Plane) findIntersection(r Ray) float64 {
	return p.Point.AddVector(*r.Origin.Negative()).DotProduct(p.Normal) / r.Direction.DotProduct(p.Normal)
}

func (p *Plane) getNormalAt(point Vector) *Vector {
	return NewVectorByVector(p.Normal)
}
