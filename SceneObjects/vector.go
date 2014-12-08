package objects

import (
	"math"
)

type Vector struct {
	x, y, z float64
}

func (this *Vector) Length() float64 {
	//return math.Sqrt(math.Pow(this.X(), 2) + math.Pow(this.Y(), 2) + math.Pow(this.Z(), 2))
	return math.Sqrt(this.X()*this.X() + this.Y()*this.Y() + this.Z()*this.Z())
}

func LimitFloat64(value, min, max float64) float64 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

func (this *Vector) Limit(min float64, max float64) *Vector {
	tx := LimitFloat64(this.X(), min, max)
	ty := LimitFloat64(this.Y(), min, max)
	tz := LimitFloat64(this.Z(), min, max)

	return NewVector(tx, ty, tz)
}

func (this *Vector) Max() float64 {
	return math.Max(this.X(), math.Max(this.Y(), this.Z()))
}

func (this *Vector) Abs() *Vector {
	return NewVector(math.Abs(this.X()), math.Abs(this.Y()), math.Abs(this.Z()))
}

func (this *Vector) Normalized() *Vector {
	var len = this.Length()
	if len == 0 {
		return NewVector(0.0, 0.0, 0.0)
	}
	return NewVector(this.X()/len, this.Y()/len, this.Z()/len)
}

func (this *Vector) Dot(o *Vector) float64 {
	return this.X()*o.X() + this.Y()*o.Y() + this.Z()*o.Z()
}

func (this *Vector) Cross(o *Vector) *Vector {
	return NewVector(this.Y()*o.Z()-this.Z()*o.Y(),
		this.Z()*o.X()-this.X()*o.Z(),
		this.X()*o.Y()-this.Y()*o.X())
}

func (this *Vector) Reflect(normal *Vector) *Vector {
	return this.Normalized().Sub(normal.Normalized().MulVal(2 * this.Normalized().Dot(normal.Normalized())))
}

func (this *Vector) AddVector(o *Vector) *Vector {
	return NewVector(this.X()+o.X(), this.Y()+o.Y(), this.Z()+o.Z())
}

func (this *Vector) AddVal(o float64) *Vector {
	return NewVector(this.X()+o, this.Y()+o, this.Z()+o)
}

func (this *Vector) Sub(o *Vector) *Vector {
	return NewVector(this.X()-o.X(), this.Y()-o.Y(), this.Z()-o.Z())
}

func (this *Vector) SubVal(o float64) *Vector {
	return NewVector(this.X()-o, this.Y()-o, this.Z()-o)
}

func (this *Vector) Mul(o *Vector) *Vector {
	return NewVector(this.X()*o.X(), this.Y()*o.Y(), this.Z()*o.Z())
}

func (this *Vector) MulVal(o float64) *Vector {
	return NewVector(this.X()*o, this.Y()*o, this.Z()*o)
}

func (this *Vector) Div(o *Vector) *Vector {
	if o.X() == 0 || o.Y() == 0 || o.Z() == 0 {
		return NewVector(0.0, 0.0, 0.0)
	}
	return NewVector(this.X()/o.X(), this.Y()/o.Y(), this.Z()/o.Z())
}

func (this *Vector) DivVal(o float64) *Vector {
	if o == 0 {
		return NewVector(0.0, 0.0, 0.0)
	}
	return NewVector(this.X()/o, this.Y()/o, this.Z()/o)
}

func (this *Vector) FindClosest(this1, this2 *Vector) *Vector {
	var dst1 = this.Sub(this1).Length()
	var dst2 = this.Sub(this2).Length()
	if dst1 < dst2 {
		return this1
	} else {
		return this2
	}
}

func NewVector(x, y, z float64) *Vector {
	tmthis := new(Vector)
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
