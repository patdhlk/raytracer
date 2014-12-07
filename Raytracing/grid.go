package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
)

type Grid struct {
	top_left     *objects.Vector
	bottom_right *objects.Vector
}

func (p *Grid) TopLeft() *objects.Vector                    { return p.top_left }
func (p *Grid) BottomRight() *objects.Vector                { return p.bottom_right }
func (p *Grid) SetTopLeft(top_left *objects.Vector)         { p.top_left = top_left }
func (p *Grid) SetBottomRight(bottom_right *objects.Vector) { p.bottom_right = bottom_right }

func NewGrid(top_left *objects.Vector, bottom_right *objects.Vector) *Grid {
	tmp := new(Grid)
	tmp.SetTopLeft(top_left)
	tmp.SetBottomRight(bottom_right)
	return tmp
}
