// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // ПОПАСТЬСЯ ДОЛЖЕН ПРАВИЛЬНЫЙ СИМВОЛ
// func isValid(s string) bool {
// 	// var s_st string = "({["
// 	// var s_end string = "]})"

// 	//ПЕРЕВОДИМ СТРОКУ В МАССИВ
// 	ss := strings.Split(s, "")

// 	//ЕСЛИ ДЛИНА МАССИВА НЕ ЧЕТНАЯ: false
// 	if len(s) % 2 != 0 {
// 		return false
// 	}

// 	var para = map[string]int {"()": 0, "{}": 0, "[]": 0,}

// 	for _, v := range ss{
// 		fmt.Print(v)
// 		if v == "(" {
// 			para["()"] = para["()"] + 1
// 		}else if v == "{" {
// 			para["{}"] = para["{}"] + 1
// 		}else if v == "[" {
// 			para["[]"] = para["[]"] + 1
// 		}
// 		fmt.Println(para)

// 		if v == ")" {
// 			if para["()"] == 0 { // или меньше 1 ???
// 				return false
// 			}else {
// 				para["()"] = para["()"] - 1
// 			}
// 		}else if v == "}" {
// 			if para["{}"] == 0 {
// 				return false
// 			}else {
// 				para["{}"] = para["{}"] - 1
// 			}
// 		}else if v == "]" {
// 			if para["[]"] == 0 {
// 				return false
// 			}else {
// 				para["[]"] = para["[]"] - 1
// 			}
// 		}
// 	}
// 	if para["()"] == 0 && para["{}"] == 0 && para["[]"] == 0 {
// 		return true
// 	}else {
// 		return false
// 	}
// }

// func main() {
// 	fmt.Println(isValid("[({{}[]()}}[]]{}"))
// 	fmt.Println(isValid("([)]"))
// }

package main
import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	var s_st string = "({["

	//ПЕРЕВОДИМ СТРОКУ В МАССИВ
	ss := strings.Split(s, "")
	//ЕСЛИ ДЛИНА МАССИВА НЕ ЧЕТНАЯ: false
	if len(ss) % 2 != 0 {
		return false
	}

	for i, v := range ss[:len(ss)-1] {
		if v == "!"{
			continue
		}else if strings.Contains(s_st, v) != true {
			return false
		}

		var para = map[string]string {"(": ")", "{": "}", "[": "]",}
		clo_par := para[v]
		ss[i] = "!"

		fmt.Println("mm", ss)

		for ii := (i + 1); ii < len(ss); ii += 2 {
			if ss[ii] == clo_par {
				ss[ii] = "!"
				fmt.Println(ss[ii])
				break
			}else {
				return false
			}
		}
	}
	return true
}


func main() {
	//fmt.Println(isValid("[({{}[]()}}[]]{}"))
	fmt.Println(isValid("()[][]"))
}