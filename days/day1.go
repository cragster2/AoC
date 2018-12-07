package days

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

func D1P1() {
	var data = utils.ReadFile("days/day1-1.txt")

	var total int = 0

	for _, freq := range data {
		value, err := strconv.Atoi(freq[1:])
		if err == nil {
			sign := freq[0]
			if sign == '+' {
				total += value
			} else {
				total -= value
			}
		}
	}
	fmt.Print("Day1 (puzzle 1): Frequency = " + strconv.Itoa(total) + "\n")
}

func D1P2() {
	var data = utils.ReadFile("days/day1-1.txt")

	var total int = 0
	var detector []int
	var freqRepeat bool = false

	detector = append(detector, 0)

	for !freqRepeat {
		for _, freq := range data {
			value, err := strconv.Atoi(freq[1:])
			if err == nil {
				sign := freq[0]
				if sign == '+' {
					total += value
				} else {
					total -= value
				}
			}
			if contains(detector, total) {
				freqRepeat = true
				break
			} else {
				detector = append(detector, total)
			}
		}
	}
	fmt.Print("Day1 (puzzle 1): Frequency repeated for " + strconv.Itoa(total) + "\n")
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
