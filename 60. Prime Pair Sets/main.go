package main

import (
	"fmt"
	"strconv"
)

const N = 100000

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	primes := [4]int{3, 7, 109, 673}
	sum := 792
	var stop bool
	for i := 675; i < N; i += 2 {
		if IsPrime(i) {
			iStr := strconv.Itoa(i)
			for k := 0; k < 4; k++ {
				kStr := strconv.Itoa(primes[k])
				rStr := iStr + kStr
				r, _ := strconv.Atoi(rStr)
				if !IsPrime(r) {
					break
				}
				rStr = kStr + iStr
				r, _ = strconv.Atoi(rStr)
				if !IsPrime(r) {
					break
				}
				stop = true
			}
			if stop {
				sum += i
				break
			}
		}
	}
	fmt.Printf("Found: %d\n", sum)
}
