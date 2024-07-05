// Написать функцию, которая получает строку
// и возвращает строку без дублирования символов
// abbcdddef -> abcdef

package main
import (
	"fmt"
)


func ReplaceDouble(s string) string {
	runes := []rune{}

	var prev rune = -1

	for _, r := range s {
		if r != prev {
			runes = append(runes, r)
			prev = r
		}
	}

	return string(runes)
}


func main() {
	fmt.Println(ReplaceDouble("abbcdddef"))
}