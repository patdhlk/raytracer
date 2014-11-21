package raytracer

//representing a point in the room
type vector struct {
	X, Y, Z int32

	// I NEED A FUCKING CONSTRUCTUR
}

//constructor
//see: https://golang.org/doc/effective_go.html#composite_literals
func NewVector(x, y, z int32) *vector {
	v := new(vector)

	//more stuff here

	return v
}

func (v *vector) Multiply(a, b vector) int32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (v *vector) Subtract(a, b vector) vector {
	var result vector
	result.X = a.X - b.X
	result.Y = a.Y - b.Y
	result.Z = a.Z - b.Z
	return result
}
