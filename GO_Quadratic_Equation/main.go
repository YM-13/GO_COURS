package main

import (
	"fmt"
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
	if ! s.Valid(eq) {
		fmt.Println("bad equation")
		return
	}

	solutions := s.Solve(eq)
	fmt.Printf("we got %d solutions\n", len(solutions))

	for _, solution := range solutions {
		fmt.Println(solution)
	}
}

// ------ solvers ----------
type FunnySolver struct {}

func (s FunnySolver) Valid(eq Equation) bool {
	return eq.A != 0
}

func (s FunnySolver) Solve(eq Equation) []float64 {
	d := math.Pow(eq.B, 2) - 4 * eq.A * eq.C
	next := math.Sqrt(d)
	res := []float64 {}
	for i := 1; i < 3; i++ {
		x := (-eq.B + next) / 2 * eq.A
		res = append(res, x)
		if d == 0.0 {
			return res
		}
		next *= -1

	}

	return res
}

func main() {
	eq := Equation{A: 1, B: -2, C: -3}

	solve(eq, FunnySolver{})
	//solve(eq, NormalSolver{})
}

// ! true == false // (1 != 2) == true
// ! false == true // (2 != 2) == false
