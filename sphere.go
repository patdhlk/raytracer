package main

import "image/color"

import "math"

type Sphere struct {
	radius   float64
	location Vector
	color    color.RGBA
}

func (this Sphere) FindIntersection(ray Ray) (float64, color.RGBA, bool) {
	d := ray.origin.AddVector(this.location.Negative())

	dDotB := d.DotProduct(ray.direction) //Solution from db
	inSqrtPart := dDotB*dDotB + this.radius*this.radius - d.DotProduct(d)

	if inSqrtPart < 0 { //Value in square root negativ -> No intersection with ball
		return 0, color.RGBA{0, 0, 0, 0}, false
	}

	//Getting nearest of 2 possible intersections
	var intersectionDistance float64

	if inSqrtPart == 0 {
		intersectionDistance = -dDotB
	} else {
		t1 := -dDotB + math.Sqrt(inSqrtPart)
		t2 := -dDotB - math.Sqrt(inSqrtPart)
		if t1 < t2 {
			intersectionDistance = t1
		} else {
			intersectionDistance = t2
		}
	}

	return intersectionDistance, this.color, true
}

func (this Sphere) GetReflectionRay(ray Ray, intersectionDistance float64) Ray {
	locationOfIntersection := ray.origin.AddVector(ray.direction.MultiplyVector(intersectionDistance))

	normal := locationOfIntersection.AddVector(this.location.Negative())
	//normal = Vector{-normal.x, -normal.y, normal.z} //Adjust x and y of normal

	reflect := CalcReflecion(ray.direction, normal)
	return NewRay(locationOfIntersection, reflect)
}
func NewSphere(location Vector, radius float64, color color.RGBA) Sphere {
	b := Sphere{radius, location, color}
	return b
}
