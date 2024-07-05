package figures

import "math"

type Point struct {
  X float64
  Y float64
}

func (p Point) DistanceTo(other Point) float64 {
  dx := other.X - p.X
  dy := other.Y - p.Y

  return math.Sqrt(dx*dx + dy*dy)
}
