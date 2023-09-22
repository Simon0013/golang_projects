package main

import (
	"fmt"
)

func main() {
	a := 100
	b := 100
	c := 800

	for true {
		if a*a+b*b == c*c {
			fmt.Printf("Found: %d; a = %d, b = %d, c = %d\n", a*b*c, a, b, c)
			break
		}
		//fmt.Printf("%d %d %d\n", a, b, c)
		if a >= c || b >= c {
			a = 250
			b++
			c = 750 - b
		} else {
			a++
			c--
		}
	}
}
