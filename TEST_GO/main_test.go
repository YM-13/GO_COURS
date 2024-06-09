package main

import "testing"

func Add(x, y int) int {
  return x + y
}

func BenchmarkOne(b *testing.B) {
    s1 := "12"

    for i := 0; i < b.N; i++ {

      res := reverse2(s1)
      _ = res

    }
}

// 2  124.3 ns/op -  29.28 ns/op
// 4  311.5 ns/op -  28.76 ns/op
// 8  609.1 ns/op -  42.92 ns/op
// 16 1270 ns/op  -  50.56 ns/op

// ~/AggravatingMediumpurpleChords$ go test -bench=. *.go
// goos: linux
// goarch: amd64
// cpu: AMD EPYC 7B12
// BenchmarkOne-6      1000000000           0.5040 ns/op
// PASS
// ok      command-line-arguments  0.622s