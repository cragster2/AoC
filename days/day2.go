package days

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func D2P1() {
	var data = utils.ReadFile("days/day2-1.txt")

	var totalPair int = 0
	var totalTriple int = 0

	for _, id := range data {
		pair, triple := checkPairTriple(id)
		totalPair += pair
		totalTriple += triple
	}
	fmt.Print("Day2 (puzzle 1): Checksum = " + strconv.Itoa(totalPair*totalTriple) + "\n")
}

func D2P2() {
	var data = utils.ReadFile("days/day2-1.txt")

	s1, s2 := compareByLetter(data)

	fmt.Print("Day1 (puzzle 1): Corresponding IDs: " + s1 + " and " + s2 + "\n")
}

func checkPairTriple(id string) (int, int) {
	pair := 0
	triple := 0

	splittedId := strings.Split(id, "")

	for _, c := range splittedId {
		total := strings.Count(id, c)
		if total == 3 {
			strings.Trim(id, c)
			triple = 1
			break
		}
	}
	for _, c := range splittedId {
		total := strings.Count(id, c)
		if total == 2 {
			pair = 1
			break
		}

	}

	return pair, triple
}

func compareByLetter(ids []string) (string, string) {

	for i, id := range ids {
		//compare to each id
		for _, id2 := range ids[i+1:] {
			total := 0
			for j, _ := range id {
				//fmt.Print(id + " " + id2 + " " + "\n")
				same := id[j] == id2[j]
				if same {
					total++
				}
				if total == len(id)-1 {
					return id, id2
				}
			}
		}
	}
	return "", ""
}
