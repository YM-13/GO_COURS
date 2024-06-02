
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
    if len(s) % 2 != 0 {
		return false
    }
	rs := []rune(s)
	var m_v = map[rune]rune {40: 41, 123: 125, 91: 93,}
	flag := 0
	count := 0

	for i, r := range(rs) {
        flag = 0
		if r == 0 {
			continue
		}else if _, ok := m_v[r]; ok == false {
			return false
		}


		for ii := i + 1; ii < len(rs); ii += 2 {
			if rs[ii] == r {
				count += 1
			}else if rs[ii] == 0 {
				return false
			}else if rs[ii] == m_v[r] {
				if count == 0 {
					fmt.Println(count)
					rs[ii] = 0
                	flag = 1
					break
				} else {
					fmt.Println("tut0000")
					count -= 1
					flag = 1
				}
			}
		}
		//fmt.Println(count)
		fmt.Println(rs)

	}
	if flag == 0 {
		fmt.Println("tut")
        return false
    }
	return true
}

func main() {
	// fmt.Println(isValid("(}]]{}"))
	// fmt.Println(isValid("[{}]"))
	fmt.Println(isValid("(([]){})")) // true
}