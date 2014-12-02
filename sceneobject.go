package main

import "image/color"

type SceneObject interface {
	FindIntersection(ray Ray) (float64, color.RGBA, bool)
	GetReflectionRay(ray Ray, intersectionDistance float64) Ray
}
