package main

import "image/png"
import "image/color"
import "os"
import "log"

var viewPoint ViewPoint
var light Light
var screen *Screen
var openingAngleOfCamera float64
var correctionOfSquareChessboard float64

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

func GetColorOfPixelInImage(ray Ray, items []SceneObject, itemBefore int, recursionDeepness int, reflectionStrength float32) (color.RGBA, bool) {

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

		distance, _, _, intersection := item.RayIntersection(ray)

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
		_, redirectedRay, color, _ := items[actualShortestInAllItems].RayIntersection(ray) //Get detailed information of intersection item

		// TO DO: Calc ray from intersection to light pos
		rayToLight := NewRay(redirectedRay.origin, light.Position.AddVector(redirectedRay.origin.Negative()))

		isShadow := false
		// is shadow?
		for _, _item := range items {
			_distance, _, _, _ := _item.RayIntersection(rayToLight)

			if _distance > 0 {
				isShadow = true
				break
			}
		}
		if isShadow {
			//fmt.Println("IsShadow %v %v", ray, redirectedRay)
		}
		// func CalcBrightness(in float64, b, L, N, V Vector, m Material, n float64) float64 {
		// TO DO: L: Einheitsvektor in Richtung der Lichtquelle
		// TO DO: N: Normale des scenceobject
		//CalcBrightness(255, ray.direction, rayToLight.direction.Normalize(), nil, redirectedRay, Polyethylene, 20)

		//Starting recursion,
		colorOfNextIntersectionItem, useThisColor := GetColorOfPixelInImage(redirectedRay, items, actualShortestInAllItems, recursionDeepness-1, reflectionStrength)
		if useThisColor == false {
			if isShadow {
				//color.B = 0
				//color.G = 0
				//color.R = 0
			}
			return color, true

		} else {
			if isShadow {
				//color.B = 0
				//color.G = 0
				//color.R = 0
			}
			return MergeColors(color, colorOfNextIntersectionItem, reflectionStrength), true
		}
	}
}

func main() {
	log.Println("Starting programm Raytracing (Stefan B., Matthias F.)")

	//Setting global variables
	screen = NewScreen(1600, 1200) //width, height
	viewPoint = ViewPoint{Vector{17, 17, -60}}
	light = Light{Vector{30, 30, 30}}
	openingAngleOfCamera = 200
	correctionOfSquareChessboard = 4

	//Local variables
	recursionDeepness := 5         //Maximum number of reflextions between items
	var reflection float32 = 0.001 //Value between 0.0 and 1.0. Should be under 0.003, if not it creates wrong colors
	items := make([]SceneObject, 0)
	nameOfOutputFile := "test.png"

	//Add Balls
	items = append(items, NewSphere(NewVector(4, 9, 14), 9.0, color.RGBA{200, 100, 0, 255}))
	items = append(items, NewSphere(NewVector(25, 9, 10), 4.0, color.RGBA{50, 25, 80, 255}))
	//items = append(items, NewSphere(NewVector(22, 8, 5), 6.0, color.RGBA{80, 0, 255, 255}))
	//items = append(items, NewSphere(NewVector(20, 20, 20), 8.0, color.RGBA{255, 255, 25, 255}))
	items = append(items, NewPlane(NewVector(0, -2, 3), NewVector(0, 8, 1), color.RGBA{0, 0, 0, 255})) //Location, direction, color

	log.Println("Screen size:", screen.width, "x", screen.height)
	log.Println("Location of Viewpoint:", viewPoint.Point.x, viewPoint.Point.y, viewPoint.Point.z)
	log.Println("Nr of items to intersect with:", len(items))
	log.Println("Recursion deepness:", recursionDeepness)

	for index, item := range items {
		log.Println("Values of item", index, ":", item)
	}

	for x := 0; x <= screen.width; x++ {
		for y := 0; y <= screen.height; y++ {
			//x is the x axis, positiv is right
			//y is the y axis, positiv is top

			r := viewPoint.CreateRayFromViewPointThroughScreen(x, y)

			colorOfPixel, _ := GetColorOfPixelInImage(r, items, -1, recursionDeepness, reflection)

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
