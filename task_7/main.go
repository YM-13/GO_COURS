package main

import "fmt"

var pln = fmt.Println
var str, new_str string
var flag bool = true

func main () {
	for flag {
		pln("Enter the string. To complete entering strings, enter 'stop'")
		fmt.Scan(&new_str)

		if new_str == "stop" {
			flag = false
			pln(len(str))
			break
		}

		if len(new_str) > len(str) {
			str = new_str
		}

	}

}