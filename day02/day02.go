package day02

// This was done in bash as follows:
// cat input.txt | xargs | sed -e 's/^+//' | bc

import (
	"fmt"

	"github.com/chazdnato/aoc/sidwtrw"
)

// SolutionOne is the first solution for the day
func SolutionOne() {
	boxIDs := sidwtrw.SliceOfStrings("day02/input.txt")
	twos, threes := 0, 0
	for _, ID := range boxIDs {
		charMap := make(map[rune]int)
		for _, c := range ID {
			charMap[c]++
		}

		for _, v := range charMap {
			if v == 2 {
				twos++
				break
			}
		}

		for _, v := range charMap {
			if v == 3 {
				threes++
				break
			}
		}
	}

	fmt.Printf("Answer: %d\n", twos*threes)
}

// SolutionTwo is the second solution for the day
func SolutionTwo() {
	boxIDs := sidwtrw.SliceOfStrings("day02/input.txt")

	matchIDs := make(map[string]int)

	for _, ID := range boxIDs {
		uniqIDs := make(map[string]string)

		// create a map of all possible IDs missing a character
		for i := range ID {
			truncatedID := ID[:i] + ID[i+1:]
			uniqIDs[truncatedID] = "aoc"
		}

		// range over all possible IDs and see if there are two that match
		for x := range uniqIDs {
			if _, ok := matchIDs[x]; ok {
				fmt.Printf("Answer: %v\n", x)
				break
			}
			matchIDs[x]++
		}
	}
}
