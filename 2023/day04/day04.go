package day04

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/chazdnato/aoc/sidwtrw"
)

func parseNumbers(numbers string) []int {
	var parsedNum []int

	for _, possibleNum := range strings.Split(numbers, " ") {
		if possibleNum == "" {
			continue
		}

		num, err := strconv.Atoi(possibleNum)
		sidwtrw.HandleError("converting gamenum", err)

		parsedNum = append(parsedNum, num)
		// fmt.Printf("'%d'\t", num)
	}

	slices.Sort(parsedNum)
	return parsedNum
}

// SolutionOne is the first solution for the day
func SolutionOne() {
	lines := sidwtrw.SliceOfStrings("day04/input.txt")
	total := 0

	for _, line := range lines {
		ordinal := 0
		hasWinning := false
		card := strings.Split(line, ":")[1]
		winning := parseNumbers(strings.Split(card, "|")[0])
		myNumbers := parseNumbers(strings.Split(card, "|")[1])
		// fmt.Printf("Winning numbers: %v\tMy numbers: %v\n", winning, myNumbers)

		for _, myNum := range myNumbers {
			if inSlice, err := sidwtrw.InSlice(myNum, winning); inSlice && err == nil {
				ordinal += 1
				hasWinning = true
				// fmt.Printf("%d is in the winning numbers\n", myNum)
			} else if err != nil {
				sidwtrw.HandleError("type mismatch in InSlice", err)
			}
		}
		// fmt.Printf("Ordinal for %s is %d\n", line, ordinal)

		if hasWinning {
			// need to remove one from the ordinal, since the first number to match
			// only counts one, and this is a naive way to do so
			total += int(math.Pow(2.0, float64(ordinal-1.0)))
		}

	}

	fmt.Printf("%d\n", total)
}
