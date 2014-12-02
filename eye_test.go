package main

import "testing"

func TestCreateRayFromViewPointThroughScreen(t *testing.T) {
	e := Eye{NewVector(0, 0, 0)}

	r := e.CreateRayFromEyeThroughScreen(0, 0, 80.0, *NewScreen(640, 480))

	if r.origin != NewVector(0, 0, 0) || r.direction != NewVector(-4.0, -3.0, 10).Normalize() {
		t.Errorf("TestCreateRayFromViewPointThroughScreen %v %v", r, NewVector(-4.0, -3.0, 10).Normalize())
	}
}
