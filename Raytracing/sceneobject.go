package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
)

type SceneObject interface {
	Intersection(*objects.Ray) (position *objects.Vector, color *objects.Vector,
		normal *objects.Vector, diffuse float64,
		specularIntensity float64, specularPower float64, reflectivity float64)
	Color() *objects.Vector
}
