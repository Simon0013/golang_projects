package main

import (
	"fmt"
	"strconv"
)

func IsPalindrome(n int) bool {
	if n < 10 {
		return true
	}
	s := strconv.Itoa(n)
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	var result int
	idx1 := 999
	idx2 := 999
	var num1, num2 int
	for idx1 > 99 {
		n := idx1 * idx2
		if IsPalindrome(n) {
			if result < n {
				result = n
				num1 = idx1
				num2 = idx2
			}
			idx1--
			idx2 = idx1
		} else if idx2 > 100 {
			idx2--
		} else {
			idx1--
			idx2 = idx1
		}
	}
	fmt.Printf("Found: %d; nums %d and %d\n", result, num1, num2)
}
