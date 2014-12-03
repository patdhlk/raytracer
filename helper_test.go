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

func TestMergeColors(t *testing.T) {
	c1 := color.RGBA{100, 200, 100, 200}
	c2 := color.RGBA{5, 10, 15, 20}

	res := MergeColors(c1, c2, 0.25)

	if res.R != 140 || res.G != 24 || res.B != 14 || res.A != 155 {
		t.Errorf("TestMergeColors 1 %v", res)
	}

	c1 = color.RGBA{100, 200, 100, 200}
	c2 = color.RGBA{50, 100, 150, 200}

	res = MergeColors(c1, c2, 0.5)

	if res.R != 75 || res.G != 150 || res.B != 125 || res.A != 200 {
		t.Errorf("TestMergeColors 2 %v", res)
	}

	c1 = color.RGBA{100, 200, 100, 200}
	c2 = color.RGBA{50, 100, 150, 200}

	res = MergeColors(c1, c2, -1)

	if res.R != 100 || res.G != 200 || res.B != 100 || res.A != 200 {
		t.Errorf("TestMergeColors 3 %v", res)
	}

	c1 = color.RGBA{100, 200, 100, 200}
	c2 = color.RGBA{50, 100, 150, 200}

	res = MergeColors(c1, c2, 1.1)

	if res.R != 100 || res.G != 200 || res.B != 100 || res.A != 200 {
		t.Errorf("TestMergeColors 3 %v", res)
	}
}
