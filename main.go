package main

import (
	"aoc/days"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string

	if len(os.Args) == 2 {
		text = os.Args[1]
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")
	}

	switch text {
	case "day1":
		days.D1P1()
		days.D1P2()
	case "day2-1":
		days.D2P1()
	case "day2-2":
		days.D2P2()
	case "day3-1":
		days.D3P1()
	case "day3-2":
		days.D3P2()
	}
}
