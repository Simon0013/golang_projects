package main

import (
	"fmt"
)

const N = 10000

func d(n int) int {
	var sum int
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	var result int
	for i := 1; i <= N; i++ {
		if d(d(i)) == i && i != d(i) {
			result += i
		}
	}
	fmt.Printf("Found: %d\n", result)
}
