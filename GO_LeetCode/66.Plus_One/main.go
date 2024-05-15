package main

import "fmt"

func plusOne(digits []int) []int {
	dg := digits
	var l_i int = len(dg) - 1
	if dg[l_i] < 9 {
		dg[l_i] = dg[l_i] + 1
		return dg
	}
	// новый вариант 14.05.24
	for i := l_i; i > -1; i-- {
		if dg[i] == 9 {
			dg[i] = 0
		} else {
			dg[i] = dg[i] + 1
			return dg
		}
	}
	res := []int {1}
	res = append(res, dg...)

	return res
}

func main() {
	var a = [] int {8, 9, 9}
	fmt.Println(plusOne(a))
}
