package main

import (
	"fmt"
)

const MONDAY = 0
const TUESDAY = 1
const WEDNESDAY = 2
const THURSDAY = 3
const FRIDAY = 4
const SATURDAY = 5
const SUNDAY = 6

func getDayWeek(start int, plus int) int {
	tmp := start + plus
	return tmp % 7
}

func IsYearVisokos(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%4 == 0 {
		return year%100 != 0
	}
	return false
}

func main() {
	countOfFirstSundays := 0
	countOfMonths := 12 * 100
	startDayOfWeek := TUESDAY //01.01.1901
	currentYear := 1901

	daysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	for i := 1; i < countOfMonths; i++ {
		startDayOfWeek = getDayWeek(startDayOfWeek, daysInMonth[(i-1)%12])
		if startDayOfWeek == SUNDAY {
			countOfFirstSundays += 1
		}
		//december was coming, Happy New Year!!
		if i%12 == 0 {
			currentYear += 1
			if IsYearVisokos(currentYear) {
				daysInMonth[1] = 29
			} else {
				daysInMonth[1] = 28
			}
		}
	}

	fmt.Printf("Found: %d\n", countOfFirstSundays)
}
