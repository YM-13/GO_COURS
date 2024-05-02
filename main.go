package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	ss "strings"
)

var pln = fmt.Println // экземпляр функции Println

var s string         // var string
var sum_nums int     // var int

var numbers []int    // slice int
var res []int        // slice int

var num int          // var int
var flag bool = true // Flag
var everige int      // var int

// THE FUNCTION RECEIVES A STRING AS INPUT AND RETURNS A STRING CLEARED OF SPACES TO THE LEFT
func data_insert(str string) []int {
	str = ss.TrimSpace(str)
	str_slice := ss.Split(str, " ")
	num_slice := make([]int, 0, len(str_slice))

	for _, n := range str_slice {
		n, err := strconv.Atoi(n)
		if err != nil {
			panic("No num")
		}
		sum_nums += n
		num_slice = append(num_slice, n)
	}
	return num_slice
}

// СЛАЙС МЭЙКЕР НАЙТИ СРЕДНЕЕ АРИФМЕТИЧЕСКОЕ 2 ОТСОРТИРОВАТЬ ЭЛЕМЕНТЫ
func slice_sorted(e_n int, args ...int) []int {
	for _, n := range args {
		if n > e_n {
			res = append(res, n)
		}
	}
	return res
}

// МИНИМАЛЬНОЕ ЧИСЛО СЛАЙСА
func min_num(args ...int) int {
	var min_n int = args[0]
	for _, n := range args[1:] {
		if n < min_n {
			min_n = n
		}
	}
	return min_n
}
func main() {
	for flag {
		pln("Enter the number. To complete entering numbers, enter 's'")
		s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if s == "s" {
			flag = false
			break
	}
	var sl_num []int = data_insert(s)
	everige = sum_nums / len(sl_num)
	pln(slice_sorted(everige, sl_num...))
	pln(min_num(sl_num...))
	}
}