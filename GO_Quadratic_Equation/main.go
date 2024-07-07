package main

import (
	"fmt"
	"math/rand"
	"math"
)

// x*x - 2x + 4 = 0
// a = 1 b = -2 c = 4

type Equation struct {
	A, B, C float64
}

type Solver interface {
	Valid(Equation) bool
	Solve(Equation) []float64
}

func solve(eq Equation, s Solver) {
	if !s.Valid(eq) {
		fmt.Println("bad equation")
		return
	}

	solutions := s.Solve(eq)
	fmt.Printf("we got %d solutions\n", len(solutions))

	for _, solution := range solutions {
		fmt.Println(solution)
	}
}

type NormalSolver struct {}

func (s NormalSolver) Valid(eq Equation) bool {
	return eq.A != 0
}

func (s NormalSolver) Solve(eq Equation) []float64 {
	d := math.Pow(eq.B, 2) - 4*eq.A*eq.C

	if d < 0 {
		return []float64{}
	}

	x1 := (-eq.B + math.Sqrt(d)) / 2 * eq.A
	x2 := (-eq.B - math.Sqrt(d)) / 2 * eq.A

	if d == 0 {
		return []float64{x1}
	} else {
		return []float64{x1, x2}
	}
}

// ------ solvers ----------
type FunnySolver struct{}

func (s FunnySolver) Valid(eq Equation) bool {
	return true
}

func (s FunnySolver) Solve(eq Equation) []float64 {
	return []float64{
		rand.Float64(),
		rand.Float64(),
	}

}

func main() {
	eq := Equation{A: 1, B: -2, C: -3}
	var solver Solver

	if rand.Intn(100) % 2 == 0 { // 50-50
		solver = FunnySolver{}
	} else {
		solver = NormalSolver{}
	}


	solve(eq, solver)
}

// ! true == false // (1 != 2) == true
// ! false == true // (2 != 2) == fals
