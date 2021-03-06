package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
)

//a light implementation
type Light struct {
	position *objects.Vector
	color    *objects.Vector
}

//creates a new instance of a light
func NewLight(position, color *objects.Vector) *Light {
	light := new(Light)
	light.SetPosition(position)
	light.SetColor(color)
	return light
}

func (this *Light) Position() *objects.Vector            { return this.position }
func (this *Light) Color() *objects.Vector               { return this.color }
func (this *Light) SetPosition(position *objects.Vector) { this.position = position }
func (this *Light) SetColor(color *objects.Vector)       { this.color = color }
