package main

type ViewPoint struct {
	Point Vector
}

func (this ViewPoint) CreateRayFromViewPointThroughScreen(x, y int) Ray {
	//Moves the viewpoint to the middle of the screen
	x = x - screen.width/2
	y = y - screen.height/2

	r := NewRay(this.Point, Vector{float64(x) / openingAngleOfCamera, float64(y) / openingAngleOfCamera, 10})
	return r
}
