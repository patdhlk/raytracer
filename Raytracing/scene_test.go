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

	if *s.View() != *objects.NewVector(7.0, 8.0, 9.0) || *s.Grid().TopLeft() != *objects.NewVector(1.0, 2.0, 3.0) || *s.Grid().BottomRight() != *objects.NewVector(1.0, 5.0, 6.0) || s.Elements() == nil || *s.Ambient() != *objects.NewVector(0.0, 0.0, 0.0) || *s.SkyColor() != *objects.NewVector(0.0, 0.0, 0.0) || s.Light() == nil {
		t.Errorf("TestNewScene %v %v %v %v %v %v %v", s.View(), s.Grid().TopLeft(), s.Grid().BottomRight(), s.Elements(), s.Ambient(), s.SkyColor(), s.Light())
	}

	g = NewGrid(objects.NewVector(1.0, 2.0, 3.0), objects.NewVector(2.0, 5.0, 6.0))
	s = NewScene(objects.NewVector(7.0, 8.0, 9.0), g)
}

func TestSceneGetSet(t *testing.T) {
	g := NewGrid(objects.NewVector(1.0, 2.0, 3.0), objects.NewVector(1.0, 5.0, 6.0))
	s := NewScene(objects.NewVector(7.0, 8.0, 9.0), g)

	if *s.View() != *objects.NewVector(7.0, 8.0, 9.0) || *s.Grid().TopLeft() != *objects.NewVector(1.0, 2.0, 3.0) || *s.Grid().BottomRight() != *objects.NewVector(1.0, 5.0, 6.0) || s.Elements() == nil || *s.Ambient() != *objects.NewVector(0.0, 0.0, 0.0) || *s.SkyColor() != *objects.NewVector(0.0, 0.0, 0.0) || s.Light() == nil {
		t.Errorf("TestSceneGetSet 1 %v %v %v %v %v %v %v", s.View(), s.Grid().TopLeft(), s.Grid().BottomRight(), s.Elements(), s.Ambient(), s.SkyColor(), s.Light())
	}

	g = NewGrid(objects.NewVector(10.0, 11.0, 12.0), objects.NewVector(10.0, 14.0, 15.0))
	s.SetGrid(g)
	s.SetView(objects.NewVector(16.0, 17.0, 18.0))
	//objects := make([]SceneObject)
	//objects = append(objects, element)
	s.SetAmbient(objects.NewVector(19.0, 20.0, 21.0))
	s.SetSkyColor(objects.NewVector(22.0, 23.0, 24.0))

	if *s.View() != *objects.NewVector(16.0, 17.0, 18.0) || *s.Grid().TopLeft() != *objects.NewVector(10.0, 11.0, 12.0) || *s.Grid().BottomRight() != *objects.NewVector(10.0, 14.0, 15.0) || s.Elements() == nil || *s.Ambient() != *objects.NewVector(19.0, 20.0, 21.0) || *s.SkyColor() != *objects.NewVector(22.0, 23.0, 24.0) {
		t.Errorf("TestSceneGetSet 2 %v %v %v %v %v %v %v", s.View(), s.Grid().TopLeft(), s.Grid().BottomRight(), s.Elements(), s.Ambient(), s.SkyColor(), s.Light())
	}
}
