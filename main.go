package raytracer

import (
	"log"
)

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

func main() {
	log.Println("rendering")
	width := 640
	height := 480

	//starts from left side of the image for each pixel to the right side of the image
	for x := 0; x < width; x++ {
		for i := 0; i < height; i++ {

		}
	}
}
