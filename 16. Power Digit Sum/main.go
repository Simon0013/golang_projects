package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var sum int
	result := math.Pow(2, 1000)
	resStr := strconv.FormatFloat(result, 'f', -1, 64)
	for _, digit := range resStr {
		sum += int(digit - '0')
	}
	fmt.Printf("Found: %d\n", sum)
}
