package raytracer

import (
	//"de/vorlesung/projekt/raytracer"
	"image"
	"image/color"
	"log"
)

func main() {
	log.Println("main started")

	var width, height int
	width = 640
	height = 480

	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{width, height}})

	O := NewVector(0, 0, 0)
	X := NewVector(1, 0, 0)
	Y := NewVector(0, 1, 0)
	//Z := NewVector(0, 0, 1)

	new_sphere_location := NewVector(1.75, -0.25, 0)

	//camera
	campos := NewVector(3, 1.5, -4)
	look_at := NewVector(0, 0, 0)
	diff_btw := NewVector(campos.x-look_at.x, campos.y-look_at.y, campos.z-look_at.z)

	camdir := diff_btw.Negative().Normalize()
	camright := Y.CrossProduct(*camdir).Negative()
	camdown := camright.CrossProduct(*camdir)
	scene_cam := Camera{*campos, *camdir, *camright, *camdown}

	white_light := color.RGBA{255, 255, 255, 255}
	pretty_green := color.RGBA{0, 255, 0, 1}
	maroon := color.RGBA{165, 74, 0, 1}
	tile := color.RGBA{255, 0, 0, 1}
	//gray := color.RGBA{142, 138, 138, 1}
	//black := color.RGBA{0, 0, 0, 1}

	//light
	light_position := NewVector(-7, 10, -10)
	scene_light := NewLight(*light_position, white_light)
	//var light_sources []Vector
	//light_sources = append(light_sources, scene_light.position)

	//scene objects
	scene_sphere := NewSphere(*O, 1, pretty_green)
	scene_sphere2 := NewSphere(*new_sphere_location, 0.5, maroon)
	scene_plane := NewPlane(*Y, *X, tile)
	var scene_objects []GraphicalObject
	scene_objects = append(scene_objects, scene_sphere, scene_sphere2, scene_plane)

	//iterating through the image
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			tempRed := 0.0
			tempGreen := 0.0
			tempBlue := 0.0

			cam_ray_origin := scene_cam.pos
			cam_ray_direction := scene_cam.dir

			cam_ray := NewRay(cam_ray_origin, cam_ray_direction)

			var intersections []float64
			for i := 0; i < len(scene_objects); i++ {

				resCount, x1, _ := scene_objects[i].findIntersection(*cam_ray)
				if resCount >= 1 {
					intersections = append(intersections, x1)
				}
			}

			var accuracy float64
			accuracy = 0.00000001
			
			index_of_winning_object := winningObjectIndex(intersections)

			if index_of_winning_object != -1 {
				// index coresponds to an object in our scene
				if intersections[index_of_winning_object] > accuracy {
					// determine the position and direction vectors at the point of intersection
					/*
						Vect intersection_position = cam_ray_origin.vectAdd(cam_ray_direction.vectMult(intersections.at(index_of_winning_object)));
						Vect intersecting_ray_direction = cam_ray_direction;
						Color intersection_color = getColorAt(intersection_position, intersecting_ray_direction, scene_objects, index_of_winning_object, light_sources, accuracy, ambientlight);
						tempRed = intersection_color.getColorRed();
						tempGreen = intersection_color.getColorGreen();
						tempBlue = intersection_color.getColorBlue();
					*/
					tempBlue = (uint8)tempRed
				}
			}
			//
			c := color.RGBA{(uint8)tempRed, (uint8)tempGreen, (uint8)tempBlue, (uint8)0}
			m.Set(x, y, c)
		}
	}
}
