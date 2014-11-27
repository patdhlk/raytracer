package raytracer

type Camera struct {
	pos   Vector
	dir   Vector
	right Vector
	down  Vector
}

func NewCamera(pos, dir, right, down Vector) *Camera {
	c := Camera{pos, dir, right, down}

	return &c
}
