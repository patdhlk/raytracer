package main

type Ray struct {
	origin    Vector
	direction Vector
}

func NewRay(origin, direction Vector) Ray {
	direction = direction.Normalize()
	ray := Ray{origin, direction}
	return ray
}
