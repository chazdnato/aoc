package day01

// This was done in bash as follows:
// cat input.txt | xargs | sed -e 's/^+//' | bc

import (
	"fmt"

	"github.com/chazdnato/aoc/sidwtrw"
)

// SolutionOne is the first solution for the day
func SolutionOne() {
	numbers := sidwtrw.SliceOfInts("day01/input.txt")
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	fmt.Printf("Answer: %d\n", sum)
}

// SolutionTwo is the second solution for the day
func SolutionTwo() {
	numbers := sidwtrw.SliceOfInts("day01/input.txt")

	highestNumber := false
	numberMap := make(map[int]int)
	sum := 0

	for {
		for _, num := range numbers {
			sum += num
			// could also do: if numberMap[sum] > 0 { ... }
			if _, ok := numberMap[sum]; ok {
				highestNumber = true
				break
			} else {
				numberMap[sum]++
			}
		}
		if highestNumber {
			break
		}
	}

	fmt.Printf("Answer: %d\n", sum)
}
