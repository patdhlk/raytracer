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
	eye         Vector
	camera      Camera
	light       Light
	sphere      Sphere
}

func NewRaytracer(w, h int, e Vector, c Camera, l Light, s Sphere) *Raytracer {
	r := Raytracer{w, h, e, c, l, s}
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
			//vector from eye to image
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

func winningObjectIndex(object_intersections []float64) int {
	var index_of_minimum_value int

	if len(object_intersections) == 0 {
		return -1
	} else if len(object_intersections) == 1 {
		if object_intersections[0] > 0 {
			return 0
		} else {
			return -1
		}
	} else {

		var max float64

		for i := range object_intersections {
			if max < object_intersections[i] {
				max = object_intersections[i]
			}
		}

		if max > 0 {
			for i := range object_intersections {
				if object_intersections[i] > 0 && object_intersections[i] <= max {
					max = object_intersections[i]
					index_of_minimum_value = i
				}
			}
			return index_of_minimum_value
		} else {
			return -1
		}
	}
}
