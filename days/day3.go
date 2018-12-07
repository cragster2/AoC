package days

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func D3P1() {

	var data = utils.ReadFile("days/data/day3-1.txt")

	var sheet [2000][2000]int
	overrides := 0

	for _, claim := range data {
		sheet, overrides = updateSheet(claim, sheet, overrides)
	}

	fmt.Print("Day2 (puzzle 1): Checksum = " + strconv.Itoa(overrides) + "\n")
}

func D3P2() {
	var data = utils.ReadFile("days/data/day3-1.txt")

	var sheet [2000][2000]int
	goodClaim := ""

	for _, claim := range data {
		sheet, _ = updateSheet(claim, sheet, 0)
	}

	for id, claim := range data {
		ok := verifySheet(claim, sheet)
		if ok {
			goodClaim = data[id]
			break
		}
	}

	fmt.Print("Day2 (puzzle 1): Checksum = " + goodClaim + "\n")
}

func updateSheet(claim string, sheet [2000][2000]int, overrides int) ([2000][2000]int, int) {

	claim = strings.Replace(claim, " @ ", "l", 1)
	claim = strings.Replace(claim, ": ", "h", 1)
	marginLeft, _ := strconv.Atoi(claim[strings.Index(claim, "l")+1 : strings.Index(claim, ",")])
	marginTop, _ := strconv.Atoi(claim[strings.Index(claim, ",")+1 : strings.Index(claim, "h")])
	length, _ := strconv.Atoi(claim[strings.Index(claim, "h")+1 : strings.Index(claim, "x")])
	height, _ := strconv.Atoi(claim[strings.Index(claim, "x")+1:])

	//fmt.Print(strconv.Itoa(marginLeft) + "," + strconv.Itoa(marginTop) + "," + strconv.Itoa(length) + "," + strconv.Itoa(height) + "\n")
	//fmt.Print(overrides)
	//fmt.Print("\n")

	for i := marginLeft; i < (marginLeft + length); i++ {
		for j := marginTop; j < (marginTop + height); j++ {
			if sheet[i][j] == 1 {
				overrides++
			}
			sheet[i][j]++
		}
	}

	return sheet, overrides
}

func verifySheet(claim string, sheet [2000][2000]int) bool {

	claim = strings.Replace(claim, " @ ", "l", 1)
	claim = strings.Replace(claim, ": ", "h", 1)
	marginLeft, _ := strconv.Atoi(claim[strings.Index(claim, "l")+1 : strings.Index(claim, ",")])
	marginTop, _ := strconv.Atoi(claim[strings.Index(claim, ",")+1 : strings.Index(claim, "h")])
	length, _ := strconv.Atoi(claim[strings.Index(claim, "h")+1 : strings.Index(claim, "x")])
	height, _ := strconv.Atoi(claim[strings.Index(claim, "x")+1:])

	//fmt.Print(strconv.Itoa(marginLeft) + "," + strconv.Itoa(marginTop) + "," + strconv.Itoa(length) + "," + strconv.Itoa(height) + "\n")
	//fmt.Print(overrides)
	//fmt.Print("\n")
	for i := marginLeft; i < (marginLeft + length); i++ {
		for j := marginTop; j < (marginTop + height); j++ {
			if sheet[i][j] != 1 {
				return false
			}
		}
	}

	return true
}
