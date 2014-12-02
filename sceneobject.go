package main

import "image/color"

type SceneObject interface {
	FindIntersection(ray Ray) (bool, float64)
	GetReflectionRay(ray Ray, intersectionDistance float64) (Ray, error)
	GetNormalAt(point Vector) Vector
	GetColor() color.RGBA
}
