package days

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

type guard struct {
	id          string
	totalAsleep int
	timeAsleep  [60]int
}

func D4P1() {

	var data = utils.ReadFile("days/data/day4-1.txt")

	fmt.Print("Day2 (puzzle 1): Checksum = " + strconv.Itoa(overrides) + "\n")
}

func D4P2() {
	var data = utils.ReadFile("days/data/day4-1.txt")

	fmt.Print("Day2 (puzzle 1): Checksum = " + goodClaim + "\n")
}

func updateAsleep(g *guard, entry string) {

}

func getGuard(guards []guard, entry string) *guard {

	return &guards[0]
}
