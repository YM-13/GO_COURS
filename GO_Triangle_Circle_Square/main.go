package main


import (
	"fmt"
	// "main/figures"
)

type GreetParams struct {
	name string
	age int
	times int
}

func greet(params GreetParams) {
	for i := 0; i < params.times; i++ {
		fmt.Printf("Hello, %s of %d y.o.!\n", params.name, params.age)
	}
}

// func main() {
// 	greet(GreetParams{
// 		times: 1,
// 		name: "Paul",
// 	})
// }


// func main1() {
// 	circle := figures.Circle{R: 10.0, X: 0, Y:0}

// 	square := figures.NewSquare(
// 		figures.Point{0,0}, figures.Point{5,5},
// 	)

// 	triangle := figures.NewTriangle(
// 		figures.Point{0,0},
// 		figures.Point{5,0},
// 		figures.Point{3,10},
// 	)

// 	Describe("circle", circle)
// 	Describe("square", square)
// 	Describe("triangle", triangle)
// }

type Figure interface {
	Perimeter() float64
	Area() float64
}

func Describe(name string, f Figure) {
	fmt.Printf("%s has p = %f and area = %f\n",
						 name, f.Perimeter(), f.Area(),
						)
}