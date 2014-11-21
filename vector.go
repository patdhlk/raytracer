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
