package scene

import "de/vorlesung/projekt/raytracer/SceneObjects"

type Light struct {
	position *objects.Vector
	color    *objects.Vector
}

func (this *Light) GetPosition() *objects.Vector         { return this.position }
func (this *Light) GetColor() *objects.Vector            { return this.color }
func (this *Light) SetPosition(position *objects.Vector) { this.position = position }
func (this *Light) SetColor(color *objects.Vector)       { this.color = color }

func NewLight(position *objects.Vector, color *objects.Vector) *Light {
	var tmp = new(Light)
	tmp.SetPosition(position)
	tmp.SetColor(color)
	return tmp
}
