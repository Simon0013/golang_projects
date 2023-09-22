package main

import (
	"fmt"
)

const N = 1000

func GetTriangle(n int) int {
	return n * (n + 1) / 2
}

func GetSquare(n int) int {
	return n * n
}

func GetPentagonal(n int) int {
	return n * (3*n - 1) / 2
}

func GetHexagonal(n int) int {
	return n * (2*n - 1)
}

func GetHeptagonal(n int) int {
	return n * (5*n - 3) / 2
}

func GetOctagonal(n int) int {
	return n * (3*n - 2)
}

func main() {
	set := [6]int{}
	//0 - triangle, 1 - square, 2 - pentagonal, 3 - hexagonal,
	//4 - heptagonal, 5 - octagonal
	found := [6]bool{}
	var stop bool
	startDigits := 10
	t := 10
	for startDigits < 100 {
		for i := 0; i < 6; i++ {
			curStart := startDigits*100 + 10
			curEnd := curStart + 89
			stop = false
			for cur := curStart; cur <= curEnd; cur++ {
				for num := 1; num < N; num++ {
					foundTri := (!found[0]) && (GetTriangle(num) == cur)
					foundSqr := (!found[1]) && (GetSquare(num) == cur)
					foundPen := (!found[2]) && (GetPentagonal(num) == cur)
					foundHex := (!found[3]) && (GetHexagonal(num) == cur)
					foundHep := (!found[4]) && (GetHeptagonal(num) == cur)
					foundOct := (!found[5]) && (GetOctagonal(num) == cur)
					if foundSqr || foundHex || foundHep || foundOct {
						startDigits = cur % 100
						set[i] = cur
						if foundTri {
							found[0] = true
						}
						if foundSqr {
							found[1] = true
						}
						if foundPen {
							found[2] = true
						}
						if foundHex {
							found[3] = true
						}
						if foundHep {
							found[4] = true
						}
						if foundOct {
							found[5] = true
						}
						stop = true
						break
					}
				}
				if stop {
					break
				}
			}
		}
		if set[5] > 0 {
			break
		} else {
			for j := 0; j < 6; j++ {
				found[j] = false
			}
			t += 1
			startDigits = t
		}
	}
	fmt.Print("Found ")
	for i := 0; i < 6; i++ {
		fmt.Printf("%d ", set[i])
	}
}
