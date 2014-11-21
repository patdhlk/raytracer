package raytracer

type camera struct {
	pos   vector
	dir   vector
	right vector
	down  vector
}

func NewCamera() *camera {
	c := new(camera)

	return c
}
