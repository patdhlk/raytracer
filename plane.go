package main

import "image/color"
import "math"

type Plane struct {
	location        Vector
	normalDirection Vector
	color           color.RGBA
}

var correctionOfSquareChessboard float64 = 4.0

func (this Plane) FindIntersection(ray Ray) (float64, color.RGBA, bool) {
	var intersectionDistance float64 = ray.origin.AddVector(this.location.Negative()).DotProduct(this.normalDirection) / ray.direction.DotProduct(this.normalDirection)

	if intersectionDistance < 0 {
		return 0, color.RGBA{0, 0, 0, 0}, false
	}

	reflectionRay := this.GetReflectionRay(ray, intersectionDistance)
	locationOfIntersection := reflectionRay.origin

	//Create chessboard by inverting plane color if x and y are both eval or odd
	chessboardVector := locationOfIntersection.AddVector(this.location.Negative())

	var returnColor color.RGBA = this.color
	if math.Abs(math.Mod(float64(int(chessboardVector.x)), 2)) == 0 && math.Abs(math.Mod(float64(int(chessboardVector.y*correctionOfSquareChessboard)), 2)) == 0 || math.Abs(math.Mod(float64(int(chessboardVector.x)), 2)) == 1 && math.Abs(math.Mod(float64(int(chessboardVector.y*correctionOfSquareChessboard)), 2)) == 1 {
		returnColor = color.RGBA{uint8(255 - this.color.R), uint8(255 - this.color.G), uint8(255 - this.color.B), uint8(255 - this.color.A)}
	}

	return intersectionDistance, returnColor, true
}

func (this Plane) GetReflectionRay(ray Ray, intersectionDistance float64) Ray {
	locationOfIntersection := ray.origin.AddVector(ray.direction.MultiplyVector(intersectionDistance))

	reflect := CalcReflecion(ray.direction, this.normalDirection)
	return NewRay(locationOfIntersection, reflect)
}

func NewPlane(location Vector, direction Vector, Color color.RGBA) Plane {
	direction = direction.Normalize()

	p := Plane{location, direction, Color}
	return p
}
