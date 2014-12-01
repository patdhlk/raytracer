package main

import "image"

type Screen struct {
	height int
	width  int
	image  *image.RGBA
}

func NewScreen(xSize, ySize int) *Screen {
	scr := Screen{ySize, xSize, image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{xSize, ySize}})}
	return &scr
}
