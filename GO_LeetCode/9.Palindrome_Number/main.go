package main

import (
    "strconv"
)
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    s := strconv.Itoa(x)
    n := len(s)
    sb := make([]byte, n, n)

    for i := 0; i < n; i++ {
        sb[i] = s[n - 1 - i]
    }

    return s == string(sb)

}
