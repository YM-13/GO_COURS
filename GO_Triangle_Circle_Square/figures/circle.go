package figures

var PI float64 = 3.14159
// CIRCLE
type Circle struct {
  R float64
  X Point
  Y Point
}

func (c Circle) Area() float64 {
  return 2 * PI * c.R
}

func (c Circle) Perimeter() float64 {
  return PI * c.R * c.R
}
