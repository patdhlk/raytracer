package objects

import (
	"math"
)

type Vector struct {
	x float64
	y float64
	z float64
}

func (p *Vector) X() float64     { return p.x }
func (p *Vector) Y() float64     { return p.y }
func (p *Vector) Z() float64     { return p.z }
func (p *Vector) SetX(x float64) { p.x = x }
func (p *Vector) SetY(y float64) { p.y = y }
func (p *Vector) SetZ(z float64) { p.z = z }

func (p *Vector) Length() float64 {
	return math.Sqrt(math.Pow(p.x, 2) + math.Pow(p.y, 2) + math.Pow(p.z, 2))
}

func LimitFloat64(value float64, min float64, max float64) float64 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

func (p *Vector) Limit(min float64, max float64) *Vector {
	var tmp_x = LimitFloat64(p.x, min, max)
	var tmp_y = LimitFloat64(p.y, min, max)
	var tmp_z = LimitFloat64(p.z, min, max)

	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) Max() float64 {
	return math.Max(p.x, math.Max(p.y, p.z))
}

func (p *Vector) Abs() *Vector {
	var tmp_x = math.Abs(p.x)
	var tmp_y = math.Abs(p.y)
	var tmp_z = math.Abs(p.z)
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) Normalized() *Vector {
	var len = p.Length()
	var tmp_x = p.x / len
	var tmp_y = p.y / len
	var tmp_z = p.z / len
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) Dot(o *Vector) float64 {
	return p.x*o.x + p.y*o.y + p.z*o.z
}

func (p *Vector) Cross(o *Vector) *Vector {
	var tmp_x = p.y*o.z - p.z*o.y
	var tmp_y = p.z*o.x - p.x*o.z
	var tmp_z = p.x*o.y - p.y*o.x
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) Reflect(normal *Vector) *Vector {
	var p_tmp = p.Normalized()
	var n_tmp = normal.Normalized()
	return p_tmp.Sub(n_tmp.MulVal(2 * p_tmp.Dot(n_tmp)))
}

func (p *Vector) Add(o *Vector) *Vector {
	var tmp_x = p.x + o.x
	var tmp_y = p.y + o.y
	var tmp_z = p.z + o.z
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) AddVal(o float64) *Vector {
	var tmp_x = p.x + o
	var tmp_y = p.y + o
	var tmp_z = p.z + o
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) Sub(o *Vector) *Vector {
	var tmp_x = p.x - o.x
	var tmp_y = p.y - o.y
	var tmp_z = p.z - o.z
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) SubVal(o float64) *Vector {
	var tmp_x = p.x - o
	var tmp_y = p.y - o
	var tmp_z = p.z - o
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) Mul(o *Vector) *Vector {
	var tmp_x = p.x * o.x
	var tmp_y = p.y * o.y
	var tmp_z = p.z * o.z
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) MulVal(o float64) *Vector {
	var tmp_x = p.x * o
	var tmp_y = p.y * o
	var tmp_z = p.z * o
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) Div(o *Vector) *Vector {
	var tmp_x = p.x / o.x
	var tmp_y = p.y / o.y
	var tmp_z = p.z / o.z
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) DivVal(o float64) *Vector {
	var tmp_x = p.x / o
	var tmp_y = p.y / o
	var tmp_z = p.z / o
	return NewVector(tmp_x, tmp_y, tmp_z)
}

func (p *Vector) FindClosest(p1 *Vector, p2 *Vector) *Vector {
	var dst1 = p.Sub(p1).Length()
	var dst2 = p.Sub(p2).Length()
	if dst1 < dst2 {
		return p1
	} else {
		return p2
	}
}

func NewVector(x float64, y float64, z float64) *Vector {
	var tmp = new(Vector)
	tmp.SetX(x)
	tmp.SetY(y)
	tmp.SetZ(z)
	return tmp
}
