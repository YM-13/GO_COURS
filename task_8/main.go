package main

import (
	"fmt"
	"strconv"
)

var pln = fmt.Println
var s string
var sum_nums int

var numbers []int
var res []int

var num int
var flag bool = true
var everige int


// СЛАЙС МЭЙКЕР 1НАЙТИ СРЕДНЕЕ АРИФМЕТИЧЕСКОЕ 2 ОТСОРТИРОВАТЬ ЭЛЕМЕНТЫ
func slice_sorted(e_n int, args ... int) {
	for _, n := range args {
		if n > e_n {
			res = append(res, n)
		}
	}
}

func min_num (args ... int) int {
	var min_n int = args[0]
	for _, n := range args[1:] {
		if n < min_n {
			min_n = n
		}
	}
	return min_n
}
func main() {
	pln("Enter the number. To complete entering numbers, enter 's'")
	for flag {
		fmt.Scan(&s)

		if s == "s" {
			flag = false
			everige = sum_nums / len(numbers)
			//НАПЕЧАТАТЬСЛАЙС С ОТСОРТИРОВАННЫМИ ЧИСЛАМИ
			slice_sorted(everige, numbers...)
			pln(res)

			//НАПЕЧАТАТЬ МИНИМВАЛЬНЫЙ ЭЛЕМЕНТ СЛАЙСА
			//pln(min_num(numbers...))
			pln("Min num of slice: ", min_num(numbers...))
			break
		}

		//Convert string in number
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, num)
		sum_nums += num
	}

}
