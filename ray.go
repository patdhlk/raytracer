package raytracer

type Ray struct {
	Origin    Vector
	Direction Vector
}

func NewRay(origin, direction Vector) *Ray {
	r := Ray{origin, direction}
	return &r
}
