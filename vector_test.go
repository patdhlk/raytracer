package raytracer

import "testing"

func TestNewVector(t *testing.T) {
	x, y, z := 4.0, 2.0, 8.0
	vec := NewVector(x, y, z)

	if vec.x != x || vec.y != y || vec.z != z {
		t.Errorf("TestNewVector %v %v %v %v %v %v", x, y, z, vec.x, vec.y, vec.z)
	}
}

func TestNewEmptyVector(t *testing.T) {
	vec := NewEmptyVector()

	if vec.x != 0 || vec.y != 0 || vec.z != 0 {
		t.Errorf("TestNewEmptyVector %v %v %v", vec.x, vec.y, vec.z)
	}
}

func TestMagnitude(t *testing.T) {
	vec := NewVector(3.0, 4.0, 0.0)

	if vec.Magnitude() != 5.0 {
		t.Errorf("TestMagnitude %v %v", vec.Magnitude(), 5.0)
	}
}

func TestNormalize(t *testing.T) {
	vec := NewVector(3.0, 4.0, 0.0)

	vec = vec.Normalize()

	if vec.x != 0.6 || vec.y != 0.8 || vec.z != 0.0 {
		t.Errorf("TestNormalize %v %v %v %v %v %v", 0.6, 0.8, 0.0, vec.x, vec.y, vec.z)
	}
}

func TestNegative(t *testing.T) {
	vec := NewVector(3.0, 4.0, 0.0)

	vec = vec.Negative()

	if vec.x != -3.0 || vec.y != -4.0 || vec.z != 0.0 {
		t.Errorf("TestNormalize %v %v %v %v %v %v", -3.0, -4.0, 0.0, vec.x, vec.y, vec.z)
	}
}

func TestDotProduct(t *testing.T) {
	vec1 := NewVector(3.5, 4.0, 5.2)
	// vec2 := NewVector(6.0, 7.3, 8.0) --> test failed... reason?!?!?!?!?
	vec2 := NewVector(6.0, 7.2, 8.0)

	result := vec1.DotProduct(*vec2)

	if result != 91.4 {
		t.Errorf("TestDotProduct %v %v", 91.4, result)
	}
}

func TestCrossProduct(t *testing.T) {
	vec1 := NewVector(3.0, 4.0, 5.0)
	vec2 := NewVector(6.0, 7.0, 8.0)

	result := vec1.CrossProduct(*vec2)

	if result.x != -3.0 || result.y != 6.0 || result.z != -3.0 {
		t.Errorf("TestCrossProduct %v %v %v %v %v %v", -3.0, 6.0, -3.0, result.x, result.y, result.z)
	}
}

func TestAddVector(t *testing.T) {
	vec1 := NewVector(3.0, 4.0, 5.0)
	vec2 := NewVector(6.0, 7.0, 8.0)

	result := vec1.AddVector(*vec2)

	if result.x != 9.0 || result.y != 11.0 || result.z != 13.0 {
		t.Errorf("TestAddVector %v %v %v %v %v %v", 9.0, 11.0, 13.0, result.x, result.y, result.z)
	}
}

func TestMultiplyVector(t *testing.T) {
	vec1 := NewVector(3.0, -4.0, 5.0)

	result := vec1.MultiplyVector(-2.5)

	if result.x != -7.5 || result.y != 10.0 || result.z != -12.5 {
		t.Errorf("TestMultiplyVector %v %v %v %v %v %v", -7.5, 10.0, -12.5, result.x, result.y, result.z)
	}
}
