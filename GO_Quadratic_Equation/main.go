package main

import (
	"fmt"
	"math/rand"
)

// x*x - 2x + 4 = 0
// a = 1 b = -2 c = 4

type Equation struct {
	A, B, C float64
}


func main() {
	eq := Equation{A: 1, B: -2, C: 4}

	solve(eq, FunnySolver{})
	//solve(eq, NormalSolver{})
}

type Solver interface {
	Valid(Equation) bool
	Solve(Equation) []float64
}

// ! true == false
// ! false == true

// (1 != 2) == true
// (2 != 2) == false


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
	// просто случайное число от 0 до 99
	return rand.Intn(100) > 50
}

func (s FunnySolver) Solve(eq Equation) []float64 {
	return []float64 {
		0.0,
		rand.Float64(),
	}
}

