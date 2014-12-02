package main

type Eye struct {
	Point Vector
}

func (this Eye) CreateRayFromEyeThroughScreen(x, y int, angle float64, screen Screen) Ray {
	//Moves the viewpoint to the middle of the screen
	x = x - screen.width/2
	y = y - screen.height/2

	r := NewRay(this.Point, Vector{float64(x) / angle, float64(y) / angle, 10})
	return r
}
