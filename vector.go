package raytracer

//representing a point in the room
type vector struct {
	X, Y, Z float64
}

//constructor
//see: https://golang.org/doc/effective_go.html#composite_literals
func NewVector(x, y, z float64) *vector {
	v := new(vector)

	//more stuff here

	return v
}

func Multiply(a, b vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func MultiplyScalar(t float64, a vector) vector {
	var result vector
	result.X = t * a.X
	result.Y = t * a.Y
	result.Z = t * a.Z
	return result
}

func Add(a, b vector) vector {
	var result vector
	result.X = a.X + b.X
	result.Y = a.Y + b.Y
	result.Z = a.Z + b.Z
	return result
}

func Subtract(a, b vector) vector {
	var result vector
	result.X = a.X - b.X
	result.Y = a.Y - b.Y
	result.Z = a.Z - b.Z
	return result
}
