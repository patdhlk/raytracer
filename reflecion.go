package raytracer

func CalcReflecion(b, N vector) vector {
	//~V=~b - 2(~b~N)~N
	return Subtract(b, MultiplyScalar(-2*Multiply(b, N), N))
}
