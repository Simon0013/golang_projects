package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

const N = 1000000

var limit = 16

func ContainsDigit(n int, digit int) (bool, error) {
	if digit > 9 {
		return false, errors.New("Argument 'digit' must be one digit")
	}
	nStr := strconv.Itoa(n)
	for _, sym := range nStr {
		if digit == int(sym-'0') {
			return true, nil
		}
	}
	return false, nil
}

func ContainsSameDigits(n1 int, n2 int) bool {
	if len(strconv.Itoa(n1)) != len(strconv.Itoa(n2)) {
		return false
	}
	for _, sym := range strconv.Itoa(n1) {
		digit := int(sym - '0')
		contains, _ := ContainsDigit(n2, digit)
		if !contains {
			return false
		}
	}
	return true
}

func GetNextNumber(n int) int {
	if (n + 1) >= limit {
		limit = limit*10 + 6
		nlen := len(strconv.Itoa(limit))
		return int(math.Pow(10, float64(nlen-1))) + 1
	}
	return n + 1
}

func main() {
	var answer int
	for i := 12; i < N; i = GetNextNumber(i) {
		if ContainsSameDigits(i, i*2) &&
			ContainsSameDigits(i*2, i*3) &&
			ContainsSameDigits(i*3, i*4) &&
			ContainsSameDigits(i*4, i*5) &&
			ContainsSameDigits(i*5, i*6) {
			answer = i
			break
		}
	}
	fmt.Printf("Found: %d\n", answer)
}
