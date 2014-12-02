package main

import (
	"errors"
	"image/color"
)

import "math"

type Sphere struct {
	radius   float64
	location Vector
	color    color.RGBA
}

func (this Sphere) FindIntersection(ray Ray) (bool, float64) {
	d := ray.origin.AddVector(this.location.Negative())

	db := d.DotProduct(ray.direction)
	discriminant := db*db + this.radius*this.radius - d.DotProduct(d)

	if discriminant < 0 { //Value in square root negativ -> No intersection with ball
		return false, 0
	}

	t1 := -db + math.Sqrt(discriminant)
	t2 := -db - math.Sqrt(discriminant)

	//Getting nearest of 2 possible intersections
	var intersectionDistance float64

	if t1 < 0 {
		intersectionDistance = t2
	} else {
		if t2 < 0 {
			intersectionDistance = t1
		} else {
			if t1 < t2 {
				intersectionDistance = t1
			} else {
				intersectionDistance = t2
			}
		}
	}

	return true, intersectionDistance
}

func (this Sphere) GetNormalAt(point Vector) Vector {
	return point.AddVector(this.location.Negative()).Normalize()
}

func (this Sphere) GetReflectionRay(ray Ray, intersectionDistance float64) (Ray, error) {
	intersect, _ := this.FindIntersection(ray)
	if intersect {
		intersection := ray.origin.AddVector(ray.direction.MultiplyVector(intersectionDistance))
		normal := intersection.AddVector(this.location.Negative())
		reflect := CalcReflecion(ray.direction, normal)
		return NewRay(intersection, reflect), nil
	}
	return NewRay(NewVector(0, 0, 0), NewVector(0, 0, 0)), errors.New("Not valid")

}

func (this Sphere) GetColor() color.RGBA {
	return this.color
}

func NewSphere(location Vector, radius float64, color color.RGBA) Sphere {
	b := Sphere{radius, location, color}
	return b
}
