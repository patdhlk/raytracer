package objects

type Plane struct {
	pos    *Vector
	normal *Vector
}

func (p *Plane) Position() *Vector        { return p.pos }
func (p *Plane) Normal() *Vector          { return p.normal }
func (p *Plane) SetPosition(pos *Vector)  { p.pos = pos }
func (p *Plane) SetNormal(normal *Vector) { p.normal = normal }

func (p *Plane) Intersection(ray *Ray) *Vector {
	var t = p.Position().Sub(ray.Origin()).Dot(p.Normal()) / ray.Direction().Dot(p.normal)
	if t <= 0 {
		return nil
	}
	return ray.AtStep(t)
}

func NewPlane(pos, normal *Vector) *Plane {
	var tmp = new(Plane)
	tmp.SetPosition(pos)
	tmp.SetNormal(normal)
	return tmp
}
