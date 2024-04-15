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

//СЛАЙС МЭЙКЕР 1НАЙТИ СРЕДНЕЕ АРИФМЕТИЧЕСКОЕ 2 ОТСОРТИРОВАТЬ ЭЛЕМЕНТЫ
func slice_sorted () {}

func main () {
	for flag {
		pln("Enter the number. To complete entering numbers, enter 'stop'")
		fmt.Scan(&s)


		if s == "stop" {
			flag = false
			//НАПЕЧАТАТЬСЛАЙС С ОТСОРТИРОВАННЫМИ ЧИСЛАМИ
			pln(res)

			//НАПЕЧАТАТЬ МИНИМВАЛЬНЫЙ ЭЛЕМЕНТ СЛАЙСА
			pln("Min num of slice: ")
			break
		}

		//Convert string in number
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, num)
		sum_nums += num
		pln(sum_nums)
	}

}