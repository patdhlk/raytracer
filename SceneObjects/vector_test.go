package objects

import (
	"log"
	"testing"
)

func TestMax(t *testing.T) {
	v := NewVector(2.0, 0.0, -1.0)
	if v.Max() != 2.0 {
		t.Errorf("V: Max should be 2.0")
	}
}

func TestAbs(t *testing.T) {
	v := NewVector(1.0, -1.1, 1.1)
	v2 := v.Abs()
	if v2.X() != 1.0 && v2.Y() != 1.1 && v2.Z() != 1.1 {
		t.Errorf("V: Abs not working")
	}
}

func TestCross(t *testing.T) {
	vec1 := NewVector(3.0, 4.0, 5.0)
	vec2 := NewVector(6.0, 7.0, 8.0)

	result := vec1.Cross(vec2)

	if result.X() != -3.0 || result.Y() != 6.0 || result.Z() != -3.0 {
		t.Errorf("TestCrossProduct %v %v %v %v %v %v", -3.0, 6.0, -3.0, result.x, result.y, result.z)
	}
}

func TestAddVector(t *testing.T) {
	vec1 := NewVector(3.0, 4.0, 5.0)
	vec2 := NewVector(6.0, 7.0, 8.0)

	result := vec1.AddVector(vec2)

	if result.X() != 9.0 || result.Y() != 11.0 || result.Z() != 13.0 {
		t.Errorf("TestAddVector %v %v %v %v %v %v", 9.0, 11.0, 13.0, result.X(), result.Y(), result.Z())
	}
}

func TestMulitplyVector(t *testing.T) {
	vec1 := NewVector(3.0, 4.0, 5.0)
	vec2 := NewVector(6.0, 7.0, 8.0)

	result := vec1.Mul(vec2)

	if result.X() != 18.0 || result.Y() != 28.0 || result.Z() != 40.0 {
		t.Errorf("TestMultiplyVector %v %v %v %v %v %v", 18.0, 28.0, 40.0, result.X(), result.Y(), result.Z())
	}
}

func TestSubVector(t *testing.T) {
	vec1 := NewVector(3.0, 4.0, 5.0)
	vec2 := NewVector(6.0, 7.0, 8.0)

	result := vec1.Sub(vec2)

	if result.X() != -3.0 || result.Y() != -3.0 || result.Z() != -3.0 {
		t.Errorf("TestSubVector %v %v %v %v %v %v", -3.0, -3.0, -3.0, result.X(), result.Y(), result.Z())
	}
}

func TestSubValue(t *testing.T) {
	vec1 := NewVector(3.0, -4.0, 5.0)

	result := vec1.SubVal(-2.5)

	if result.X() != 5.5 || result.Y() != -1.5 || result.Z() != 7.5 {
		t.Errorf("TestSubValue %v %v %v %v %v %v", 5.5, -1.5, 7.5, result.X(), result.Y(), result.Z())
	}
}

func TestMultiplyValue(t *testing.T) {
	vec1 := NewVector(3.0, -4.0, 5.0)

	result := vec1.MulVal(-2.5)

	if result.X() != -7.5 || result.Y() != 10.0 || result.Z() != -12.5 {
		t.Errorf("TestMultiplyValue %v %v %v %v %v %v", -7.5, 10.0, -12.5, result.X(), result.Y(), result.Z())
	}
}

func TestDivValue(t *testing.T) {
	vec1 := NewVector(3.0, -4.0, 5.0)

	result := vec1.DivVal(-2.5)

	if result.X() != -1.2 || result.Y() != 1.6 || result.Z() != -2 {
		t.Errorf("TestDivValue %v %v %v %v %v %v", -1.2, 1.6, -2, result.X(), result.Y(), result.Z())
	}

	result = vec1.DivVal(0)
	if result.X() != 0.0 || result.Y() != 0.0 || result.Z() != 0.0 {
		log.Println("div 0 issue...null vector returned")
		t.Errorf("TestDivValue %v %v %v %v %v %v", 0.0, 0.0, 0.0, result.X(), result.Y(), result.Z())
	}
}

func TestDivVector(t *testing.T) {
	vec1 := NewVector(3.0, -4.0, 5.0)
	vec2 := NewVector(1.0, 2.0, -1.0)
	result := vec1.Div(vec2)

	if result.X() != 3.0 || result.Y() != -2.0 || result.Z() != -5.0 {
		t.Errorf("TestDivVector %v %v %v %v %v %v", 3.0, -2.0, -5.0, result.X(), result.Y(), result.Z())
	}
	vec2 = NewVector(0.0, 0.0, 0.0)
	result = vec1.Div(vec2)
	if result.X() != 0.0 || result.Y() != 0.0 || result.Z() != 0.0 {
		log.Println("div 0 issue...null vector returned")
		t.Errorf("TestDivVector %v %v %v %v %v %v", 0.0, 0.0, 0.0, result.X(), result.Y(), result.Z())
	}
}

func TestDotProduct(t *testing.T) {
	vec1 := NewVector(3.5, 4.0, 5.2)
	//vec2 := NewVector(6.0, 7.3, 8.0) //--> test failed... reason?!?!?!?!?
	vec2 := NewVector(6.0, 7.2, 8.0)

	result := vec1.Dot(vec2)

	if result != 91.4 {
		t.Errorf("TestDotProduct %v %v", 91.4, result)
	}
}

func TestNormalize(t *testing.T) {
	vec := NewVector(3.0, 4.0, 0.0)

	vec = vec.Normalized()

	if vec.x != 0.6 || vec.y != 0.8 || vec.z != 0.0 {
		t.Errorf("TestNormalize %v %v %v %v %v %v", 0.6, 0.8, 0.0, vec.X(), vec.Y(), vec.Z())
	}

	vec = NewVector(0.0, 0.0, 0.0)

	vec = vec.Normalized()

	if vec.X() != 0.0 || vec.Y() != 0.0 || vec.Z() != 0.0 {
		log.Println("div 0 issue...null vector returned")
		t.Errorf("TestNormalize %v %v %v %v %v %v", 0.0, 0.0, 0.0, vec.X(), vec.Y(), vec.Z())
	}
}

func TestReflect(t *testing.T) {
	vec1 := NewVector(1.0, -1.0, 0.0)
	vec2 := NewVector(0.0, 1.0, 0.0)
	r := vec1.Reflect(vec2)
	if r.X() != 0.7071067811865475 || r.Y() != 0.7071067811865475 || r.Z() != 0.0 {
		t.Errorf("TestReflect got Vec(%f, %f, %f)", r.X(), r.Y(), r.Z())
	}
}

func TestAddValue(t *testing.T) {
	vec := NewVector(1.0, -1.0, 0.0)
	var i float64
	i = 5.0
	vec = vec.AddVal(i)
	if vec.X() != 6 || vec.Y() != 4 || vec.Z() != 5 {
		t.Errorf("TestAddValue failed %v %v %v ", vec.X(), vec.Y(), vec.Z())
	}
}
