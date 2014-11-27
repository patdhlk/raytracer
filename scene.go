package raytracer

type Scene struct {
	eye    Vector
	camera Camera
	light  Light
	sphere Sphere
}

func NewScene(e Vector, c Camera, l Light, s Sphere) *Scene {
	scence := Scene{e, c, l, s}
	return &scence
}
