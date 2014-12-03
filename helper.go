package main

import "image/color"

func RemoveTransparency(c *color.RGBA) {
	c.A = 255
}

func MergeColors(c1, c2 color.RGBA, factor float32) color.RGBA {
	if factor < 0 || factor > 1 {
		return c1
	}
	r1, g1, b1, a1 := c1.RGBA()
	r2, g2, b2, a2 := c2.RGBA()
	return color.RGBA{
		uint8(float32(r1)*(1-factor)) + uint8(float32(r2)*factor),
		uint8(float32(g1)*(1-factor)) + uint8(float32(g2)*factor),
		uint8(float32(b1)*(1-factor)) + uint8(float32(b2)*factor),
		uint8(float32(a1)*(1-factor)) + uint8(float32(a2)*factor)}
}
