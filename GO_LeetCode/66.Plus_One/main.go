package main

//import "fmt"

func plusOne(digits []int) []int {
    var last_index int
    var last_new_dig int
	var res [] int

	var time_var int = 1

    last_index = len(digits) - 1
    last_new_dig = digits[last_index] + 1
    if last_new_dig < 10 {
        digits[last_index] = last_new_dig
        return digits
    } else {
		res = make([]int, 0, len(digits))
        digits[last_index] = 0
		// [9, 9, 9]
        for i := last_index - 1; i > -1; i-- {
			new_d +=
			res = append(res, digits[i] + time_var)
        }

    }
    return res
}
