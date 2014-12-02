package main

import "image/color"
import "math"

type Plane struct {
	location        Vector
	normalDirection Vector
	color           color.RGBA
}

func (this Plane) RayIntersection(ray Ray) (float64, Ray, color.RGBA, bool) {
	var intersectionDistance float64 = ray.origin.AddVector(this.location.Negative()).DotProduct(this.normalDirection) / ray.direction.DotProduct(this.normalDirection)

	if intersectionDistance < 0 {
		return 0, ray, color.RGBA{0, 0, 0, 0}, false
	}

	reflectionRay, locationOfIntersection := this.CalculateReflectionRay(ray, intersectionDistance)

	//Create chessboard by inverting plane color if x and y are both eval or odd
	chessboardVector := locationOfIntersection.AddVector(this.location.Negative())

	var returnColor color.RGBA = this.color
	if math.Abs(math.Mod(float64(int(chessboardVector.x)), 2)) == 0 && math.Abs(math.Mod(float64(int(chessboardVector.y*correctionOfSquareChessboard)), 2)) == 0 || math.Abs(math.Mod(float64(int(chessboardVector.x)), 2)) == 1 && math.Abs(math.Mod(float64(int(chessboardVector.y*correctionOfSquareChessboard)), 2)) == 1 {
		returnColor = color.RGBA{uint8(255 - this.color.R), uint8(255 - this.color.G), uint8(255 - this.color.B), uint8(255 - this.color.A)}
	}

	return intersectionDistance, reflectionRay, returnColor, true
}

func (this Plane) CalculateReflectionRay(ray Ray, intersectionDistance float64) (Ray, Vector) {
	locationOfIntersection := ray.origin.AddVector(ray.direction.MultiplyVector(intersectionDistance))

	bN := ray.direction.DotProduct(this.normalDirection)
	directionOfReflectionRay := ray.direction.AddVector(this.normalDirection.MultiplyVector(2 * bN).Negative())

	return NewRay(locationOfIntersection, directionOfReflectionRay), locationOfIntersection
}

func NewPlane(location Vector, direction Vector, Color color.RGBA) Plane {
	direction = direction.Normalize()

	p := Plane{location, direction, Color}
	return p
}
