package figures

import "math"

var RK float64 = 0.5

// TRIANGLE
type Triangle struct {
  a Point
  b Point
  c Point
}

func NewTriangle(a, b, c Point) Triangle {
  return Triangle{a: a, b:b, c:c}
}

func (t Triangle) Perimeter() float64 {
  p := t.a.DistanceTo(t.b)
  p += t.b.DistanceTo(t.c)
  p += t.c.DistanceTo(t.a)
  return p
}

func (t Triangle) Sides() (float64, float64, float64) {
  return t.a.DistanceTo(t.b),
    t.b.DistanceTo(t.c),
    t.c.DistanceTo(t.a)
}

func (t Triangle) Area() float64 {
  // a, b, c - длины сторон, p - полупериметр
  a, b, c := t.Sides()
  p := float64(a + b + c) / 2.0

  x := p * (p - a) * (p - b) * (p - c)

  return math.Sqrt(x)
}
