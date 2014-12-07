package objects

import (
	"math"
)

type Vector struct {
	x float64
	y float64
	z float64
}

func (this *Vector) Length() float64 {
	return math.Sqrt(math.Pow(this.X(), 2) + math.Pow(this.Y(), 2) + math.Pow(this.Z(), 2))
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

func (this *Vector) Limit(min float64, max float64) *Vector {
	var tmthis_x = LimitFloat64(this.X(), min, max)
	var tmthis_y = LimitFloat64(this.Y(), min, max)
	var tmthis_z = LimitFloat64(this.Z(), min, max)

	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) Max() float64 {
	return math.Max(this.X(), math.Max(this.Y(), this.Z()))
}

func (this *Vector) Abs() *Vector {
	var tmthis_x = math.Abs(this.X())
	var tmthis_y = math.Abs(this.Y())
	var tmthis_z = math.Abs(this.Z())
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) Normalized() *Vector {
	var len = this.Length()
	var tmthis_x = this.X() / len
	var tmthis_y = this.Y() / len
	var tmthis_z = this.Z() / len
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) Dot(o *Vector) float64 {
	return this.X()*o.X() + this.Y()*o.Y() + this.Z()*o.Z()
}

func (this *Vector) Cross(o *Vector) *Vector {
	tmthis_x := this.Y()*o.Z() - this.Z()*o.Y()
	tmthis_y := this.Z()*o.X() - this.X()*o.Z()
	tmthis_z := this.X()*o.Y() - this.Y()*o.X()
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) Reflect(normal *Vector) *Vector {
	var this_tmthis = this.Normalized()
	var n_tmthis = normal.Normalized()
	return this_tmthis.Sub(n_tmthis.MulVal(2 * this_tmthis.Dot(n_tmthis)))
}

func (this *Vector) Add(o *Vector) *Vector {
	var tmthis_x = this.X() + o.X()
	var tmthis_y = this.Y() + o.Y()
	var tmthis_z = this.Z() + o.Z()
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) AddVal(o float64) *Vector {
	var tmthis_x = this.X() + o
	var tmthis_y = this.Y() + o
	var tmthis_z = this.Z() + o
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) Sub(o *Vector) *Vector {
	var tmthis_x = this.X() - o.X()
	var tmthis_y = this.Y() - o.Y()
	var tmthis_z = this.Z() - o.Z()
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) SubVal(o float64) *Vector {
	var tmthis_x = this.X() - o
	var tmthis_y = this.Y() - o
	var tmthis_z = this.Z() - o
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) Mul(o *Vector) *Vector {
	var tmthis_x = this.X() * o.X()
	var tmthis_y = this.Y() * o.Y()
	var tmthis_z = this.Z() * o.Z()
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) MulVal(o float64) *Vector {
	var tmthis_x = this.X() * o
	var tmthis_y = this.Y() * o
	var tmthis_z = this.Z() * o
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) Div(o *Vector) *Vector {
	var tmthis_x = this.X() / o.X()
	var tmthis_y = this.Y() / o.Y()
	var tmthis_z = this.Z() / o.Z()
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) DivVal(o float64) *Vector {
	var tmthis_x = this.X() / o
	var tmthis_y = this.Y() / o
	var tmthis_z = this.Z() / o
	return NewVector(tmthis_x, tmthis_y, tmthis_z)
}

func (this *Vector) FindClosest(this1 *Vector, this2 *Vector) *Vector {
	var dst1 = this.Sub(this1).Length()
	var dst2 = this.Sub(this2).Length()
	if dst1 < dst2 {
		return this1
	} else {
		return this2
	}
}

func NewVector(x float64, y float64, z float64) *Vector {
	var tmthis = new(Vector)
	tmthis.SetX(x)
	tmthis.SetY(y)
	tmthis.SetZ(z)
	return tmthis
}

func (this *Vector) X() float64     { return this.x }
func (this *Vector) Y() float64     { return this.y }
func (this *Vector) Z() float64     { return this.z }
func (this *Vector) SetX(x float64) { this.x = x }
func (this *Vector) SetY(y float64) { this.y = y }
func (this *Vector) SetZ(z float64) { this.z = z }
