package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	main()
}

func TestRaytracing(t *testing.T) {
	raytracing("unittest.png")
}

func TestGetColorOfPixelInImage(t *testing.T) {
	recursionDeepness := 5
	var reflection float32 = 0.001
	s := NewDefaultScene()

	r := s.eye.CreateRayFromEyeThroughScreen(0, 0, 200, *NewScreen(640, 480))

	colorOfPixel, res := GetColorOfPixelInImage(r, s.light, s.objects, -1, recursionDeepness, reflection)

	if res != true {
		t.Errorf("TestGetColorOfPixelInImage %v %v", colorOfPixel, res)
	}
}
