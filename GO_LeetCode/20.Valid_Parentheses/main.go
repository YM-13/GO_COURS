
// 	//ЕСЛИ ДЛИНА МАССИВА НЕ ЧЕТНАЯ: false


// func main() {
// 	fmt.Println(isValid("[({{}[]()}}[]]{}"))
// 	fmt.Println(isValid("([)]"))
// }

package main
import (
	"fmt"
)

func pair(p rune) rune {
	if p == '(' {
		return ')'
	} else if p == '[' {
		return ']'
	} else if p == '{' {
		return '}'
	}
	return ' '
}

func check(s string) bool {
	stack := []rune{} // обычный слайс

	for _, r := range s {
		if r == '(' || r == '[' || r == '{' {
			stack = append(stack, r)
		} else if r == ')' || r == '}' || r == ']' {
			if len(stack) == 0 {
				return false
			} else {
				if stack[len(stack)-1] != pair(r) {
					return false
				}
				stack = stack[0 : len(stack)-1] // выкинули последний элем.
			}
		}
	}

	return len(stack) == 0
}

func main() {
	// fmt.Println(isValid("(}]]{}"))
	// fmt.Println(isValid("[{}]"))
	fmt.Println(check("[({])}")) // false
}