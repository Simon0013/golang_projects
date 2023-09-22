package main

import "fmt"

func IsPrime(n int64) bool {
	if n == 1 {
		return false
	}
	var i int64
	for i = 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var current int64
	for i := 1; i <= 10001; i++ {
		for j := current + 1; true; j++ {
			if IsPrime(j) {
				current = j
				break
			}
		}
	}

	fmt.Printf("Found: %d\n", current)
}
