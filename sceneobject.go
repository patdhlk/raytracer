package main

import "image/color"

type SceneObject interface {
	RayIntersection(ray Ray) (float64, Ray, color.RGBA, bool)
}
