package scene

import "de/vorlesung/projekt/raytracer/SceneObjects"

type Light struct {
	position *objects.Vector
	color    *objects.Vector
}

func (p *Light) GetPosition() *objects.Vector         { return p.position }
func (p *Light) GetColor() *objects.Vector            { return p.color }
func (p *Light) SetPosition(position *objects.Vector) { p.position = position }
func (p *Light) SetColor(color *objects.Vector)       { p.color = color }

func NewLight(position *objects.Vector, color *objects.Vector) *Light {
	var tmp = new(Light)
	tmp.SetPosition(position)
	tmp.SetColor(color)
	return tmp
}
