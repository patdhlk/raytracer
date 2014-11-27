package raytracer

import "testing"

func TestNewCamera(t *testing.T) {
	c := NewCamera(*NewVector(1, 2, 3), *NewVector(4, 5, 6), *NewVector(7, 8, 9), *NewVector(10, 11, 12))

	v := c.pos
	if v.x != 1 || v.y != 2 || v.z != 3 {
		t.Errorf("TestNewCamera c.pos %v %v %v %v %v %v", 1, 2, 3, v.x, v.y, v.z)
	}
	v = c.dir
	if v.x != 4 || v.y != 5 || v.z != 6 {
		t.Errorf("TestNewCamera c.dir %v %v %v %v %v %v", 4, 5, 6, v.x, v.y, v.z)
	}
	v = c.right
	if v.x != 7 || v.y != 8 || v.z != 9 {
		t.Errorf("TestNewCamera c.right %v %v %v %v %v %v", 7, 8, 9, v.x, v.y, v.z)
	}
	v = c.down
	if v.x != 10 || v.y != 11 || v.z != 12 {
		t.Errorf("TestNewCamera c.down %v %v %v %v %v %v", 10, 11, 12, v.x, v.y, v.z)
	}
}
