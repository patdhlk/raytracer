package scene

import (
	"de/vorlesung/projekt/raytracer/Helper"
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
)

type Surface struct {
	plane             *objects.Plane
	color             *objects.Vector
	diffuse           float64
	specularIntensity float64
	specularPower     float64
	reflectivity      float64
}

func (p *Surface) Intersection(line *objects.Ray) (position,
	color, normal *objects.Vector, diffuse,
	specularIntensity, specularPower, reflectivity float64) {

	position = p.plane.Intersection(line)
	if position == nil {
		return
	}
	color = p.color.Mul(p.getTextureColor(position))
	normal = p.Plane().Normal()
	diffuse = p.Diffuse()
	specularIntensity = p.SpecularIntensity()
	specularPower = p.SpecularPower()
	reflectivity = p.Reflectivity()
	return
}

func (p *Surface) getTextureColor(position *objects.Vector) *objects.Vector {
	position = position.Abs()
	var intVal1 = Helper.Round(position.X(), 0)
	var intVal2 = Helper.Round(position.Y(), 0)
	var intVal3 = Helper.Round(position.Z(), 0)
	if int(intVal1+intVal2+intVal3)%2 == 0 {
		return objects.NewVector(0.0, 0.0, 0.0)
	}
	return objects.NewVector(1.0, 1.0, 1.0)
}

func NewSurface(plane *objects.Plane, color *objects.Vector, diffuse,
	specularIntensity, specularPower,
	reflectivity float64) *Surface {
	var tmp = new(Surface)
	tmp.SetPlane(plane)
	tmp.SetColor(color)
	tmp.SetDiffuse(diffuse)
	tmp.SetSpecularIntensity(specularIntensity)
	tmp.SetSpecularPower(specularPower)
	tmp.SetReflectivity(reflectivity)
	return tmp
}

func (p *Surface) Plane() *objects.Plane          { return p.plane }
func (p *Surface) Color() *objects.Vector         { return p.color }
func (p *Surface) Diffuse() float64               { return p.diffuse }
func (p *Surface) SpecularIntensity() float64     { return p.specularIntensity }
func (p *Surface) SpecularPower() float64         { return p.specularPower }
func (p *Surface) Reflectivity() float64          { return p.reflectivity }
func (p *Surface) SetPlane(plane *objects.Plane)  { p.plane = plane }
func (p *Surface) SetColor(color *objects.Vector) { p.color = color }
func (p *Surface) SetDiffuse(diffuse float64)     { p.diffuse = diffuse }
func (p *Surface) SetSpecularIntensity(specularIntensity float64) {
	p.specularIntensity = specularIntensity
}
func (p *Surface) SetSpecularPower(specularPower float64) { p.specularPower = specularPower }
func (p *Surface) SetReflectivity(reflectivity float64)   { p.reflectivity = reflectivity }
