package main

import (
	"fmt"
	"strconv"
)

func main() {
	var temperature string
	var n_32 int64 = 32
	var n_5 int64 = 5
	var n_9 int64 = 9

	fmt.Println("Enter temperature in degrees Fahrenheit")
	fmt.Scan(&temperature)
	t_F, err := strconv.Atoi(temperature)
	if err != nil {
		panic(err)
	}

	res := (int64(t_F) - n_32) * n_5 / n_9

	fmt.Println("Temperature in degrees Celsius: ", res)
}
