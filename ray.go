package raytracer

type Ray struct {
	Origin    Vector
	Direction Vector
}

func NewRay(origin, direction Vector) *Ray {
	r := new(Ray)
	r.Origin = origin
	r.Direction = direction
	return r
}
