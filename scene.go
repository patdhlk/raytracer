package main

import "image/color"

type Scene struct {
	eye     Eye
	light   Light
	objects []SceneObject
}

func NewScene(eye Eye, light Light, objects []SceneObject) *Scene {
	sc := Scene{eye, light, objects}
	return &sc
}

func NewDefaultScene() *Scene {
	eye := Eye{Vector{17, 17, -60}}
	light := Light{Vector{30, 30, 30}}

	objects := make([]SceneObject, 0)

	objects = append(objects, NewSphere(NewVector(4, 9, 14), 9.0, color.RGBA{200, 100, 0, 255}))
	objects = append(objects, NewSphere(NewVector(25, 9, 10), 4.0, color.RGBA{50, 25, 80, 255}))
	objects = append(objects, NewPlane(NewVector(0, -2, 3), NewVector(0, 8, 1), color.RGBA{0, 0, 0, 255}))

	sc := Scene{eye, light, objects}
	return &sc
}
