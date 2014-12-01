package main

type Light struct {
	Position Vector
}

func NewLight(position Vector) Light {
	l := Light{position}
	return l
}
