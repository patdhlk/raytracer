package raytracer

import (
	"log"
)

type Raytracer struct {
	imagewidth  int
	imageheight int
	camera      Camera
	light       Light
	sphere      Sphere
}

func NewRaytracer(w, h int, c Camera, l Light, s Sphere) *Raytracer {
	r := Raytracer{w, h, c, l, s}
	return &r
}

// TODO pseudocode for basic ray tracer
// :for each screen pixel
// :generate a ray from the camera to the pixel
// :intersect the ray with all objects in the scene
// :for the closest intersection
// :determine the material
// :for each light in the scene
// :generate a ray from the intersection to the light
// :         if not obstructed: Apply illumination
// :spawn secondary rays (e.g., reflection, refraction)
// :combine results
func (r *Raytracer) RunRaytracer() {
	log.Println("start raytracing")

	for x := 0; x < r.imagewidth; x++ {
		for i := 0; i < r.imageheight; i++ {

		}
	}
}
