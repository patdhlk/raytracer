package main

import (
	"errors"
	"image/color"
)

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
	returnColor := color.RGBA{255, 255, 255, 0}

	return intersectionDistance, returnColor, true
}

func (this Plane) GetReflectionRay(ray Ray, intersectionDistance float64) (Ray, error) {
	_, _, intersect := this.FindIntersection(ray)
	if intersect {
		intersection := ray.origin.AddVector(ray.direction.MultiplyVector(intersectionDistance))
		reflect := CalcReflecion(ray.direction, this.normalDirection)
		return NewRay(intersection, reflect), nil
	}
	return NewRay(NewVector(0, 0, 0), NewVector(0, 0, 0)), errors.New("Not valid")
}

func (this Plane) GetNormalAt(point Vector) Vector {
	return this.normalDirection
}

func NewPlane(location Vector, direction Vector, Color color.RGBA) Plane {
	direction = direction.Normalize()
	p := Plane{location, direction, Color}
	return p
}
