package figures

import (
	"math"
)

// SQUARE
type Square struct {
	//A float64
	// d2= (х2— х1)2+ (y2— y1)2
	// Сумма квадратов диагонали прямоугольника равны сумме квадратов сторон:
	// 2d2 = 2a2 + 2b2

  a Point // topleft
  b Point // bottomright
}

// func findeDiagonals(x1, y1, x2, y2 float64) (float64, float64) {
// 	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)), math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
// }

// type Person struct {
//   Name string
// }

// func foo() {
//   John := Person{ Name: "John" }
// }

func NewSquare(p1, p2 Point) Square {
	instance := Square{
		a: p1,
		b: p2,
	}
	return instance
}

// func foo() {
// 	s := Square{}
// 	// s.a == Point{}
// 	// s.b == Point{}
// 	// zero-point = Point{vX: 0.0, Y; 0.0 }
// 	// s == Square {
// 	// a: Point{ 0, 0 },
// 	// b: Point{ 0, 0 },
// 	//}

// 	s = Square{
// 		// a == Point{ 0, 0 }
// 		b: Point{ X: 10, Y: 25.1 },
// 	}

// 	_ = s
// }

func (s Square) Side() float64 {
  diag := s.a.DistanceTo(s.b)
  return diag / math.Sqrt(2.0)
}

func (s Square) Perimeter() float64 {
	// dx := math.Abs(s.X1 - s.X3)
	// dy := math.Abs(s.Y1 - s.Y3)
	// return 2 * (dx + dy)
  diag := s.a.DistanceTo(s.b)
  return 4 * diag / math.Sqrt(2.0)
}

func (s Square) Area() float64 {
	return s.Side() * s.Side()
}

// 0, 0 -------> x
// |
// |   1,1 ---------- 5,1
// |    |              |
// v    |              |
// y   1,3 -----------5,3
