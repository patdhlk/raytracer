package raytracer

func Multiply(a, b vector) int32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Subtract(a, b vector) vector {
	var result vector
	result.X = a.X - b.X
	result.Y = a.Y - b.Y
	result.Z = a.Z - b.Z
	return result
}
