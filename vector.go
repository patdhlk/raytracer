package raytracer

import "math"

type Vector struct {
	x, y, z float64
}

func NewVector(x, y, z float64) *Vector {
	v := new(Vector)
	v.x = x
	v.y = y
	v.z = z
	return v
}

func NewEmptyVector() *Vector {
	v := new(Vector)
	v.x = 0
	v.y = 0
	v.z = 0
	return v
}

// func (v *vector) GetVectorX() float64 {
// 	return v.x
// }

// func (v *vector) GetVectorY() float64 {
// 	return v.y
// }

// func (v *vector) GetVectorZ() float64 {
// 	return v.z
// }

func (v *Vector) Magnitude() float64 {
	l := math.Sqrt(float64((v.x * v.x) + (v.y * v.y) + (v.z * v.z)))
	return float64(l)
}

func (v *Vector) Normalize() *Vector {
	return NewVector(v.x/v.Magnitude(), v.y/v.Magnitude(), v.z/v.Magnitude())
}

func (v *Vector) Negative() *Vector {
	return NewVector(-v.x, -v.y, -v.z)
}

func (v *Vector) DotProduct(vec Vector) float64 {
	return v.x*vec.x + v.y*vec.y + v.z*vec.z
}

func (v *Vector) CrossProduct(vec Vector) *Vector {
	return NewVector(v.y*vec.z-v.z*vec.y,
		v.z*vec.x-v.x*vec.z,
		v.x*vec.y+v.y*vec.x)
}

func (v *Vector) AddVector(vec Vector) *Vector {
	return NewVector(v.x+vec.x, v.y+vec.y, v.z+vec.z)
}

func (v *Vector) MultiplyVector(scalar float64) *Vector {
	return NewVector(v.x*scalar, v.y*scalar, v.z*scalar)
}
