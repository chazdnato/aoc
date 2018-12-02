package day02

// This was done in bash as follows:
// cat input.txt | xargs | sed -e 's/^+//' | bc

import (
	"bufio"
	"fmt"
	"os"
)

// HandleError is a generic error handler
func HandleError(location string, err error) {
	if err != nil {
		fmt.Printf("ERR: %s\t %v\n", location, err)
		os.Exit(2)
	}
}

// SliceOfStrings returns a slice of strings from the input file
func SliceOfStrings() []string {
	file, err := os.Open("day02/input.txt")
	HandleError("Opening file", err)
	defer file.Close()

	s := bufio.NewScanner(file)

	var strings []string

	for s.Scan() {
		word := s.Text()
		strings = append(strings, word)
	}
	return strings
}

// SolutionOne is the first solution for the day
func SolutionOne() {
	boxIDs := SliceOfStrings()
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
	boxIDs := SliceOfStrings()

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
