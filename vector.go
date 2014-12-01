package main

import "math"

type Vector struct {
	x float64
	y float64
	z float64
}

func NewVector(x, y, z float64) Vector {
	v := Vector{x, y, z}
	return v
}

func NewVectorEmpty() Vector {
	v := Vector{0, 0, 0}
	return v
}

func NewVectorByVector(vec Vector) Vector {
	v := Vector{vec.x, vec.y, vec.z}
	return v
}

func (this Vector) Negative() Vector {
	return Vector{-this.x, -this.y, -this.z}
}

func (this Vector) MultiplyVector(m float64) Vector {
	return Vector{this.x * m, this.y * m, this.z * m}
}

func (this Vector) AddVector(v Vector) Vector {
	return Vector{this.x + v.x, this.y + v.y, this.z + v.z}
}

func (this Vector) Normalize() Vector {
	n := this.Magnitude()
	if this == NewVectorEmpty() {
		return NewVectorEmpty()
	}
	return Vector{this.x / n, this.y / n, this.z / n}
}

func (this Vector) DotProduct(v Vector) float64 {
	return this.x*v.x + this.y*v.y + this.z*v.z
}

func (this Vector) Magnitude() float64 {
	return math.Sqrt(this.x*this.x + this.y*this.y + this.z*this.z)
}

func (this Vector) CrossProduct(v Vector) Vector {
	return Vector{this.y*v.z - this.z*v.y, this.z*v.x - this.x*v.z, this.x*v.y - this.y*v.x}
}
