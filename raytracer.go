package raytracer

import (
	"image"
	"image/png"
	"log"
	"os"
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
//the core algorithm
func (r *Raytracer) RunRaytracer() {
	log.Println(log.Ltime, " start raytracing")
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{r.imagewidth, r.imageheight}})
	log.Println(log.Ltime, " image section created")

	for x := 0; x < r.imagewidth; x++ {
		for i := 0; i < r.imageheight; i++ {

		}
	}

	//create image file
	w, _ := os.Create("rederedImage.png")
	//png encoding
	if err := png.Encode(w, m); err != nil {
		log.Fatal(log.Ltime, " image creation error", err)
		os.Exit(1)
	}
	log.Println(log.Ltime, " finished rendering")
}
