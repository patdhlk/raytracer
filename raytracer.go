package main

import (
	"de/vorlesung/projekt/raytracer/Helper"
	"de/vorlesung/projekt/raytracer/Raytracing"
	"de/vorlesung/projekt/raytracer/SceneObjects"
	"log"
	"os"
)

func main() {
	h := new(Helper.Helper)
	width := 640
	height := 480
	filename := "out.png"

	f, _ := os.Create("logfile")
	log.SetOutput(f)
	log.Println("################################")
	log.Println("start rendering")

	sphere1 := objects.NewSphere(objects.NewVector(0.0, 0.0, 1.0), 1.0)
	sphere2 := objects.NewSphere(objects.NewVector(-2.5, 0.0, -0.75), 1.0)
	plane := objects.NewPlane(objects.NewVector(0.0, -1.0, 0.0), objects.NewVector(0.0, 1.0, 0.0))

	currentScene := scene.NewScene(objects.NewVector(4.0, 0.5, 0.0), scene.NewGrid(objects.NewVector(2.0, 1.00, -1.0), objects.NewVector(2.0, -0.50, 1.0)))
	currentScene.AddElement(scene.SceneObject(scene.NewBall(sphere1, objects.NewVector(0.54, 0.60, 0.90), 0.9, 4.0, 30.0, 0.125)))
	currentScene.AddElement(scene.SceneObject(scene.NewBall(sphere2, objects.NewVector(1.0, 2.0, 0.0), 0.9, 4.0, 30.0, 0.125)))
	currentScene.AddElement(scene.SceneObject(scene.NewSurface(plane, objects.NewVector(1.0, 1.0, 1.0), 1.0, 1.0, 8.0, 0.05)))
	currentScene.SetAmbient(objects.NewVector(0.25, 0.25, 0.3))
	currentScene.SetLight(scene.NewLight(objects.NewVector(1.0, 4.0, 0.5), objects.NewVector(1.0, 1.0, 1.0)))
	currentScene.SetSkyColor(objects.NewVector(0.85, 0.85, 0.95))

	i := currentScene.Render(width, height, 1)
	h.ImageWriter(filename, i)
}
