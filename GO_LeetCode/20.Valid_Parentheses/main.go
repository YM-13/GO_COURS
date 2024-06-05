
// 	//ЕСЛИ ДЛИНА МАССИВА НЕ ЧЕТНАЯ: false


// func main() {
// 	fmt.Println(isValid("[({{}[]()}}[]]{}"))
// 	fmt.Println(isValid("([)]"))
// }

package main
import (
	"fmt"
)

func isValid(s string) bool {
    if len(s) % 2 != 0 || len(s) < 2 {
		return false
    }

	rs := []rune(s)
	var m_v = map[rune]rune {40: 41, 123: 125, 91: 93,}
	//var prev int
	count := 0

	for i := 0; i < len(rs) - 1; i++ {
		if rs[i] == 0 {
			continue
		}
		//prev = i
		for ii := i + 1; ii < len(rs); ii += 2 {
			if rs[ii] == 0 {
				continue
			}else if m_v[rs[i]] == rs[ii] {// rs[i] - key from map
				// ПРОВЕРИТЬ ПРОХОДОМ В ОБРАТНОМ ПОРЯДКЕ НЕТ ЛИ ОТКРІВАЮЩЕЙ СКОБКИ
				// ЕСЛИ ЕСТЬ, ТО "0" ПРИСВОИТЬ ЕЙ
				for b := ii - 1; -1 < b; b -= 2 {
					if rs[b] == rs[i] {
						rs[b] = 0
						rs[ii] = 0
						count += 2
						break
					}
				}
			}
		}
	}
	fmt.Println(rs)
	if count == len(rs) {
		return true
	}else {
		return false
	}


}

func main() {
	// fmt.Println(isValid("(}]]{}"))
	// fmt.Println(isValid("[{}]"))
	fmt.Println(isValid("[({])}")) // false
}