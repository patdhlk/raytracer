package raytracer

type Camera struct {
	pos   Vector
	dir   Vector
	right Vector
	down  Vector
}

func NewCamera() *Camera {
	c := new(Camera)

	return c
}
