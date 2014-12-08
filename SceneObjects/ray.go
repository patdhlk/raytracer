package objects

type Ray struct {
	origin    *Vector
	direction *Vector
}

func (this *Ray) Origin() *Vector                { return this.origin }
func (this *Ray) Direction() *Vector             { return this.direction }
func (this *Ray) SetOrigin(origin *Vector)       { this.origin = origin }
func (this *Ray) SetDirection(direction *Vector) { this.direction = direction }

func (this *Ray) AtStep(count float64) *Vector {
	return this.Origin().AddVector(this.Direction().MulVal(count))
}

func NewRay(origin, direction *Vector) *Ray {
	ray := new(Ray)
	ray.SetOrigin(origin)
	ray.SetDirection(direction.Normalized())
	return ray
}
