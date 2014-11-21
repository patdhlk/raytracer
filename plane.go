package raytracer

type Plane struct {
	e     Vector
	n     Vector
	color int64
}

func (p *Plane) findIntersection(r Ray) float64 {
	return p.e.AddVector(*r.Origin.Negative()).DotProduct(p.n) / r.Direction.DotProduct(p.n)
}
