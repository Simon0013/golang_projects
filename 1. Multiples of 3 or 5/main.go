package main

import "fmt"

const N = 1000

func main() {
	var sum int

	for i := 0; i < N; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Printf("Found sum: %d\n", sum)
}
