package raytracer

import (
	"image/color"
)

type Light struct {
	Position Vector
	Color    color.RGBA
}

func NewLight(position Vector, color color.RGBA) *Light {
	l := Light{position, color}
	return &l
}
