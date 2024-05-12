package main

import "fmt"

func plusOne(digits []int) []int {
    var last_index int
    var last_new_dig int
	var res [] int
	var new_d int
	var time_var int = 0

    last_index = len(digits) - 1
    last_new_dig = digits[last_index] + 1

    if last_new_dig < 10 {
        digits[last_index] = last_new_dig
        return digits

    } else {
		time_var = 1
		// новый слайс длина старого + 1 = [0, 0, 0, ...]. Для случая с 10 в новом слайсе "0" уже есть
		res = make([]int, len(digits) + 1)

        for i := last_index - 1; i > -1; i-- {
			if time_var == 1 {
				new_d = digits[i] + 1
				if new_d < 10 {
					time_var = 0
					res[i + 1] = new_d
					res1 := digits[: i]
					res2 := res[i + 1 :]
					res = append(res1, res2...)
					return res
				}else { // new_d == 10
					res[i + 1] = 0
				}


        	}

    	}
		res[0] = 1
    return res
	}
}

func main() {
	var a = [] int {8}
	fmt.Println(plusOne(a))
}