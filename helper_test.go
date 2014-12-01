package main

import (
	"image/color"
	"testing"
)

func TestRemoveTransparency(t *testing.T) {
	c := color.RGBA{200, 100, 0, 0}

	RemoveTransparency(&c)

	if c.A != 255 {
		t.Errorf("TestRemoveTransparency %v", c.A)
	}
}
