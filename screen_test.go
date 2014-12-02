package main

import "testing"

func TestNewScreen(t *testing.T) {
	s := NewScreen(480, 640)

	if s.height != 640 || s.width != 480 || s.image.Rect.Max.X != 480 || s.image.Rect.Max.Y != 640 {
		t.Errorf("TestNewScreen %v %v %v %v", s.height, s.width, s.image.Rect.Max.X, s.image.Rect.Max.Y)
	}
}
