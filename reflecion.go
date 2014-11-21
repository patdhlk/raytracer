package raytracer

func CalcReflecion(b, N Vector) *Vector {
	//~V=~b - 2(~b~N)~N
	return b.AddVector(*N.MultiplyVector(-2 * b.DotProduct(N)))
}
