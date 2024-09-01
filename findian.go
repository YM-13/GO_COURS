package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkString(s string) bool {
	lett := "ian"

	if s[0] == lett[0] && s[len(s)-1] == lett[2] {
		for i := 1; i < len(s)-1; i++ {
			if s[i] == lett[1] {
				return true
			}
		}

	}
	return false

}
func main() {
	var s string
	fmt.Printf("Enter a string: ")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	s = in.Text()
	s = strings.ToLower(s)

	if len(s) < 3 || !checkString(s) {
		fmt.Printf("Not Found!")
	} else {
		fmt.Println("Found!")
	}
}
