package main

import "image/png"

import "image/color"
import "os"
import "log"

/**
 * GetColorImagePixel
 *
 * Gets the color of a Pixel in Image. If there is a intersection with a item, it gets the
 * redirected ray and calls GetColorImagePixel again. So you can see reflections on the items.
 *
 * @param ray Ray to cross the items
 * @param items All items in the screen
 * @param itemBefore Item which was crossed before, is needed in recursion case
 * @param recursionDeepness How often reflextions are allowed
 * @param reflectionStrength How strong the reflection should be
 * @return color.RGBA Calculated color if there was a intersection
 * @return bool Says, if the returned color should be used or not.Happens, if the ray goes nowhere
 */
func GetColorOfPixelInImage(ray Ray, light Light, items []SceneObject, itemBefore int, recursionDeepness int, reflectionStrength float32) (color.RGBA, bool) {

	if recursionDeepness == 0 {
		return color.RGBA{0, 0, 0, 0}, false
	}

	var minDistanceOfAllItems float64 = 99999999999
	actualShortestInAllItems := 0
	intersectionWithAnyItem := false

outer:
	for index, item := range items { //Find nearest item
		if itemBefore == index { //Avoids crossing the same item again, could happen because of rounding errors
			continue outer
		}

		distance, _, intersection := item.FindIntersection(ray)

		if intersection == false { //No intersection with item
			continue outer
		}

		if distance < 0 { //If distance < 0, intersection behind viewpoint -> uninteresting intersection
			continue outer
		}

		if distance < minDistanceOfAllItems {
			intersectionWithAnyItem = true
			minDistanceOfAllItems = distance
			actualShortestInAllItems = index
		}

	}

	if intersectionWithAnyItem == false {
		return color.RGBA{0, 0, 0, 0}, false
	} else {
		_distance, color, _ := items[actualShortestInAllItems].FindIntersection(ray) //Get detailed information of intersection item

		redirectedRay, _ := items[actualShortestInAllItems].GetReflectionRay(ray, _distance)

		rayToLight := NewRay(redirectedRay.origin, light.Position.AddVector(redirectedRay.origin.Negative()))

		isShadow := false
		for _index, _item := range items {
			if itemBefore != _index {
				_distance, _, _ := _item.FindIntersection(rayToLight)

				if _distance < 0 {
					isShadow = true
					break
				}
			}
		}

		p := Polyethylene{}
		normal := items[actualShortestInAllItems].GetNormalAt(redirectedRay.origin)
		brightness := CalcBrightness(255, ray.direction, rayToLight.direction.Normalize(), normal, redirectedRay.direction, p)

		colorOfNextIntersectionItem, useThisColor := GetColorOfPixelInImage(redirectedRay, light, items, actualShortestInAllItems, recursionDeepness-1, reflectionStrength)
		if useThisColor == false {
			if isShadow {
				color.B = 0
				color.G = 0
				color.R = 0

			}
			color.A = uint8(brightness)
			return color, true

		} else {
			if isShadow {
				color.B = 0
				color.G = 0
				color.R = 0
			}

			color = MergeColors(color, colorOfNextIntersectionItem, reflectionStrength)
			color.A = uint8(brightness)
			return color, true
		}
	}
}

func raytracing(file string) {
	log.Println("Start raytracing")

	screen := NewScreen(1600, 1200)
	angle := 200.0

	//Local variables
	recursionDeepness := 5
	var reflection float32 = 0.001

	nameOfOutputFile := file

	scene := NewDefaultScene()

	for x := 0; x <= screen.width; x++ {
		for y := 0; y <= screen.height; y++ {
			r := scene.eye.CreateRayFromEyeThroughScreen(x, y, angle, *screen)

			colorOfPixel, _ := GetColorOfPixelInImage(r, scene.light, scene.objects, -1, recursionDeepness, reflection)

			RemoveTransparency(&colorOfPixel)
			screen.image.SetRGBA(x, y, colorOfPixel)
		}
	}

	log.Println("Creating output file", nameOfOutputFile)
	w, _ := os.Create(nameOfOutputFile)
	if err := png.Encode(w, screen.image); err != nil {
		log.Println("Error writing image on Disk")
		os.Exit(1)
	}
}

func main() {
	raytracing("default.png")
}
