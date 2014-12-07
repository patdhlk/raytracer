package scene

import "de/vorlesung/projekt/raytracer/SceneObjects"

type Ball struct {
	sphere            *objects.Sphere
	color             *objects.Vector
	diffuse           float64
	specularIntensity float64
	specularPower     float64
	reflectivity      float64
}

func (p *Ball) Intersection(line *objects.Ray) (position *objects.Vector, color *objects.Vector,
	normal *objects.Vector, diffuse float64,
	specularIntensity float64, specularPower float64, reflectivity float64) {

	position = p.sphere.Intersection(line)
	if position == nil {
		return
	}
	color = p.color
	normal = position.Sub(p.sphere.GetPosition())
	diffuse = p.diffuse
	specularIntensity = p.specularIntensity
	specularPower = p.specularPower
	reflectivity = p.reflectivity
	return
}

func NewBall(sphere *objects.Sphere, color *objects.Vector,
	diffuse float64, specularIntensity float64, specularPower float64,
	reflectivity float64) *Ball {
	tmp := new(Ball)
	tmp.SetSphere(sphere)
	tmp.SetColor(color)
	tmp.SetDiffuse(diffuse)
	tmp.SetSpecularIntensity(specularIntensity)
	tmp.SetSpecularPower(specularPower)
	tmp.SetReflectivity(reflectivity)
	return tmp
}

func (p *Ball) Sphere() *objects.Sphere          { return p.sphere }
func (p *Ball) Color() *objects.Vector           { return p.color }
func (p *Ball) Diffuse() float64                 { return p.diffuse }
func (p *Ball) SpecularIntensity() float64       { return p.specularIntensity }
func (p *Ball) SpecularPower() float64           { return p.specularPower }
func (p *Ball) Reflectivity() float64            { return p.reflectivity }
func (p *Ball) SetSphere(sphere *objects.Sphere) { p.sphere = sphere }
func (p *Ball) SetColor(color *objects.Vector)   { p.color = color }
func (p *Ball) SetDiffuse(diffuse float64)       { p.diffuse = diffuse }
func (p *Ball) SetSpecularIntensity(specularIntensity float64) {
	p.specularIntensity = specularIntensity
}
func (p *Ball) SetSpecularPower(specularPower float64) { p.specularPower = specularPower }
func (p *Ball) SetReflectivity(reflectivity float64)   { p.reflectivity = reflectivity }
