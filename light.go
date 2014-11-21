package raytracer

type Light struct {
	Position Vector
	Color    int64
}

func NewLight(position Vector, color int64) *Light {
	l := new(Light)
	l.Position = position
	l.Color = color
	return l
}
