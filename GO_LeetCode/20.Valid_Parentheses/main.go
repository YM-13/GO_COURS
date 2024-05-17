package main

import (
	"fmt"
	"strings"
)

// ПОПАСТЬСЯ ДОЛЖЕН ПРАВИЛЬНЫЙ СИМВОЛ


func isValid(s string) bool {
	var s_st string = "({["
	var s_end string = "]})"

	ss := strings.Split(s, "")
	fmt.Println(ss)
	if len(s) % 2 != 0 {
		fmt.Println("---")
		return false
	}else if strings.Contains(s_st, ss[0]) != true || strings.Contains(s_end, ss[len(s) - 1]) != true {
		fmt.Println("---")
		return false
	}
	fmt.Println("+++")
	return true
}

func main() {
	isValid("[({{}[]()}}[]]{}")
}