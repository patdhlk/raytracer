package objects

import (
	"math"
)

type Sphere struct {
	position *Vector
	radius   float64
}

func (p *Sphere) Position() *Vector            { return p.position }
func (p *Sphere) Radius() float64              { return p.radius }
func (p *Sphere) SetPosition(position *Vector) { p.position = position }
func (p *Sphere) SetRadius(radius float64)     { p.radius = radius }

func (p *Sphere) Intersection(ray *Ray) *Vector {
	var d = ray.Origin().Sub(p.Position())
	var db = ray.Direction().Dot(d)
	var subSqrt = math.Pow(db, 2) + math.Pow(p.Radius(), 2) - math.Pow(d.Length(), 2)
	if subSqrt < 0 {
		return nil
	}
	var sqrt = math.Sqrt(subSqrt)
	var t1 = -db + sqrt
	var t2 = -db - sqrt
	if t1 < 0 && t2 < 0 {
		return nil
	}
	var jumps = math.Min(t1, t2)
	if t1 < 0 {
		jumps = t2
	}
	if t2 < 0 {
		jumps = t1
	}
	return ray.AtStep(jumps)
}

func NewSphere(position *Vector, radius float64) *Sphere {
	var tmp = new(Sphere)
	tmp.SetPosition(position)
	tmp.SetRadius(radius)
	return tmp
}
