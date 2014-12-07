package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
)

type SceneObject interface {
	Intersection(*objects.Ray) (position, color,
		normal *objects.Vector, diffuse,
		specularIntensity, specularPower, reflectivity float64)
	Color() *objects.Vector
}
