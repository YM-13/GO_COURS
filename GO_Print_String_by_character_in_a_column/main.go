// Написать функцию, которая получает строку
// и печатает ее по символам в столбик

package main

import (
	"fmt"
)

// snake_case <- do not use!
// camelCase
// PascalCase
// kebab-case

func PrintStr(s string) {
	// for _, r := range s {
	// 	fmt.Println(string(r))
	// }

	for _, r := range []rune(s) {
		fmt.Printf("%s\n", string(r))
	}
}

func main() {
	PrintStr("Urban")
}