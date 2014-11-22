package raytracer

type GraphicalObject interface {
	findIntersection(r Ray) (int32, float64, float64)
	getNormalAt(point Vector) *Vector
}
