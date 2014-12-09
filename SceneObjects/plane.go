package objects

//basic plane implementation
type Plane struct {
	pos    *Vector
	normal *Vector
}

//calculates the intersection between the plane and one ray
//returns the position of intersection, a vector
func (p *Plane) Intersection(ray *Ray) *Vector {
	t := p.Position().Sub(ray.Origin()).Dot(p.Normal()) / ray.Direction().Dot(p.normal)
	if t <= 0 {
		return nil
	}
	return ray.AtStep(t)
}

//creates a new instance of a plane
func NewPlane(pos, normal *Vector) *Plane {
	var tmp = new(Plane)
	tmp.SetPosition(pos)
	tmp.SetNormal(normal)
	return tmp
}

//Getters and Setters

//returns the position of the plane
func (p *Plane) Position() *Vector        { return p.pos }
func (p *Plane) Normal() *Vector          { return p.normal }
func (p *Plane) SetPosition(pos *Vector)  { p.pos = pos }
func (p *Plane) SetNormal(normal *Vector) { p.normal = normal }
