package main

import "fmt"

const N = 100

func main() {
	tmp := int64(float64((1 + N)) / 2 * float64(N))
	sqrOfSum := tmp * tmp

	var sumSqrs int64
	for i := 1; i <= N; i++ {
		sumSqrs += int64(i * i)
	}

	fmt.Printf("Found: %d - %d = %d\n", sqrOfSum, sumSqrs, sqrOfSum-sumSqrs)
}
