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

func (this Plane) FindIntersection(ray Ray) (bool, float64) {
	var intersectionDistance float64 = this.location.AddVector(ray.origin.Negative()).DotProduct(this.normalDirection) / ray.direction.DotProduct(this.normalDirection)

	if intersectionDistance < 0 {
		return false, 0
	}

	return true, intersectionDistance
}

func (this Plane) GetReflectionRay(ray Ray, intersectionDistance float64) (Ray, error) {
	intersect, _ := this.FindIntersection(ray)
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

func (this Plane) GetColor() color.RGBA {
	return this.color
}

func NewPlane(location Vector, direction Vector, Color color.RGBA) Plane {
	direction = direction.Normalize()
	p := Plane{location, direction, Color}
	return p
}
