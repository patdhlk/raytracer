package raytracer

type Light struct {
	Position Vector
	Color    int64
}

func NewLight(position Vector, color int64) *Light {
	l := Light{position, color}
	return &l
}
