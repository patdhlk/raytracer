package scene

import (
	objects "de/vorlesung/projekt/raytracer/SceneObjects"
	"testing"
)

func TestNewScene(t *testing.T) {
	defer func() {
		e := recover()
		if e == nil {
			t.Errorf("TestNewScene error is nil")
		}
	}()
	g := NewGrid(objects.NewVector(1.0, 2.0, 3.0), objects.NewVector(1.0, 5.0, 6.0))
	s := NewScene(objects.NewVector(7.0, 8.0, 9.0), g)

	if *s.eye != *objects.NewVector(7.0, 8.0, 9.0) || *s.grid.top_left != *objects.NewVector(1.0, 2.0, 3.0) || *s.grid.bottom_right != *objects.NewVector(1.0, 5.0, 6.0) || s.elements == nil || *s.ambient != *objects.NewVector(0.0, 0.0, 0.0) || *s.skyColor != *objects.NewVector(0.0, 0.0, 0.0) || s.light == nil {
		t.Errorf("TestNewScene %v %v %v %v %v %v %v", s.eye, s.grid.top_left, s.grid.bottom_right, s.elements, s.ambient, s.skyColor, s.light)
	}

	g = NewGrid(objects.NewVector(1.0, 2.0, 3.0), objects.NewVector(2.0, 5.0, 6.0))
	s = NewScene(objects.NewVector(7.0, 8.0, 9.0), g)
}

func TestSceneGetSet(t *testing.T) {
	g := NewGrid(objects.NewVector(1.0, 2.0, 3.0), objects.NewVector(1.0, 5.0, 6.0))
	s := NewScene(objects.NewVector(7.0, 8.0, 9.0), g)

	if *s.Eye() != *objects.NewVector(7.0, 8.0, 9.0) || *s.Grid().top_left != *objects.NewVector(1.0, 2.0, 3.0) || *s.Grid().bottom_right != *objects.NewVector(1.0, 5.0, 6.0) || s.Elements() == nil || *s.Ambient() != *objects.NewVector(0.0, 0.0, 0.0) || *s.SkyColor() != *objects.NewVector(0.0, 0.0, 0.0) || s.light == nil {
		t.Errorf("TestSceneGetSet 1 %v %v %v %v %v %v %v", s.Eye(), s.Grid().top_left, s.Grid().bottom_right, s.Elements(), s.Ambient(), s.SkyColor(), s.Light())
	}

	g = NewGrid(objects.NewVector(10.0, 11.0, 12.0), objects.NewVector(10.0, 14.0, 15.0))
	s.SetGrid(g)
	s.SetEye(objects.NewVector(16.0, 17.0, 18.0))
	//objects := make([]SceneObject)
	//objects = append(objects, element)
	s.SetAmbient(objects.NewVector(19.0, 20.0, 21.0))
	s.SetSkyColor(objects.NewVector(22.0, 23.0, 24.0))

	if *s.Eye() != *objects.NewVector(16.0, 17.0, 18.0) || *s.Grid().top_left != *objects.NewVector(10.0, 11.0, 12.0) || *s.Grid().bottom_right != *objects.NewVector(10.0, 14.0, 15.0) || s.Elements() == nil || *s.Ambient() != *objects.NewVector(19.0, 20.0, 21.0) || *s.SkyColor() != *objects.NewVector(22.0, 23.0, 24.0) {
		t.Errorf("TestSceneGetSet 2 %v %v %v %v %v %v %v", s.Eye(), s.Grid().top_left, s.Grid().bottom_right, s.Elements(), s.Ambient(), s.SkyColor(), s.Light())
	}
}
