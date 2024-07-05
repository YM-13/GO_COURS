// Написать функцию reverse, которая возвращает строку
// задом-наперед "hello world" -> "dlrow olleh"

package main

import (
	"fmt"
)

func reverse(s string) string {
	var res string

	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}

	return res
}


func reverse2(s string) string {
	n := len(s) // len in bytes
	bs := make([]byte, n, n)

	for i := 0; i < len(s); i++ {
		// _ = bs[len(s) - 1]

		bs[i] = s[len(s) - 1 - i]
	}

	return string(bs)
}

func main() {
	//PrintStr("hello мир")
	// fmt.Println(string(97))
	//fmt.Println(ReplaceDouble("abbcdddef"))
	fmt.Println(reverse2("rtut"))
}