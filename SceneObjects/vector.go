//Package objects implements the scene objects: vector, ray, plane and sphere
//vector = Vektor
//ray = Lichtstrahl
//plane = Ebene
//sphere = Kugel
package objects

import (
	"math"
)

//vector implementation
type Vector struct {
	x, y, z float64
}

//the length of a vector
func (this *Vector) Length() float64 {
	//return math.Sqrt(math.Pow(this.X(), 2) + math.Pow(this.Y(), 2) + math.Pow(this.Z(), 2))
	return math.Sqrt(this.X()*this.X() + this.Y()*this.Y() + this.Z()*this.Z())
}

//just a helper function for the limit func
func helper(value, min, max float64) float64 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

//limit of a vector
func (this *Vector) Limit(min float64, max float64) *Vector {
	tx := helper(this.X(), min, max)
	ty := helper(this.Y(), min, max)
	tz := helper(this.Z(), min, max)

	return NewVector(tx, ty, tz)
}

//the maximum value (x, y or z) of a vector
func (this *Vector) Max() float64 {
	return math.Max(this.X(), math.Max(this.Y(), this.Z()))
}

//the absolute value of a vector
func (this *Vector) Abs() *Vector {
	return NewVector(math.Abs(this.X()), math.Abs(this.Y()), math.Abs(this.Z()))
}

//function to normalize a vector
func (this *Vector) Normalized() *Vector {
	var len = this.Length()
	if len == 0 {
		return NewVector(0.0, 0.0, 0.0)
	}
	return NewVector(this.X()/len, this.Y()/len, this.Z()/len)
}

//returns the dotproduct of two vector
func (this *Vector) Dot(o *Vector) float64 {
	return this.X()*o.X() + this.Y()*o.Y() + this.Z()*o.Z()
}

//returns the cross product two vectors
func (this *Vector) Cross(o *Vector) *Vector {
	return NewVector(this.Y()*o.Z()-this.Z()*o.Y(),
		this.Z()*o.X()-this.X()*o.Z(),
		this.X()*o.Y()-this.Y()*o.X())
}

//the vector reflection method
func (this *Vector) Reflect(normal *Vector) *Vector {
	return this.Normalized().Sub(normal.Normalized().MulVal(2 * this.Normalized().Dot(normal.Normalized())))
}

//adds two vectors
func (this *Vector) AddVector(o *Vector) *Vector {
	return NewVector(this.X()+o.X(), this.Y()+o.Y(), this.Z()+o.Z())
}

// func (this Vector) CheapAddVector(o *Vector) Vector {
// 	v := new(Vector)
// 	v.SetX(this.X() + o.X())
// 	v.SetY(this.Y() + o.Y())
// 	v.SetZ(this.Z() + o.Z())
// 	return *v
// }

//adds a value to a vector
//returns new vector
func (this *Vector) AddVal(o float64) *Vector {
	return NewVector(this.X()+o, this.Y()+o, this.Z()+o)
}

//subtract o vector from another vector (this)
//returns new vector
func (this *Vector) Sub(o *Vector) *Vector {
	return NewVector(this.X()-o.X(), this.Y()-o.Y(), this.Z()-o.Z())
}

//subtracts a value from a vector
//returns new vector
func (this *Vector) SubVal(o float64) *Vector {
	return NewVector(this.X()-o, this.Y()-o, this.Z()-o)
}

//multiply two vectors
//returns the new vector
func (this *Vector) Mul(o *Vector) *Vector {
	return NewVector(this.X()*o.X(), this.Y()*o.Y(), this.Z()*o.Z())
}

//multiply a value with one vector
//returns the new vector
func (this *Vector) MulVal(o float64) *Vector {
	return NewVector(this.X()*o, this.Y()*o, this.Z()*o)
}

//divide two vectors
//returns the new vector
func (this *Vector) Div(o *Vector) *Vector {
	if o.X() == 0 || o.Y() == 0 || o.Z() == 0 {
		return NewVector(0.0, 0.0, 0.0)
	}
	return NewVector(this.X()/o.X(), this.Y()/o.Y(), this.Z()/o.Z())
}

//devide a vector through a value
//returns the new vector
func (this *Vector) DivVal(o float64) *Vector {
	if o == 0 {
		return NewVector(0.0, 0.0, 0.0)
	}
	return NewVector(this.X()/o, this.Y()/o, this.Z()/o)
}

//find the closest vector, out of two, to the vector,
//returns the nearer vector
func (this *Vector) FindClosest(this1, this2 *Vector) *Vector {
	var dst1 = this.Sub(this1).Length()
	var dst2 = this.Sub(this2).Length()
	if dst1 < dst2 {
		return this1
	} else {
		return this2
	}
}

//creates a new instance of a vector
func NewVector(x, y, z float64) *Vector {
	tmthis := new(Vector)
	tmthis.SetX(x)
	tmthis.SetY(y)
	tmthis.SetZ(z)
	return tmthis
}

//Getters and Setters

//returns the x coordinate
func (this *Vector) X() float64 { return this.x }

//return the y coordinate
func (this *Vector) Y() float64 { return this.y }

//return the z coordinate
func (this *Vector) Z() float64 { return this.z }

//sets the x coordinate
func (this *Vector) SetX(x float64) { this.x = x }

//sets the y coordinate
func (this *Vector) SetY(y float64) { this.y = y }

//sets the z coordinate
func (this *Vector) SetZ(z float64) { this.z = z }
