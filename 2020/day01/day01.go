package day01

import (
	"fmt"
	"sort"

	"github.com/chazdnato/aoc/sidwtrw"
)

// SolutionOne is the first solution for the day
func SolutionOne() {
	numbers := sidwtrw.SliceOfInts("day01/input.txt")
	product := 0

	sort.Ints(numbers)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] > 2020 {
				continue
			}
			if numbers[i]+numbers[j] == 2020 {
				fmt.Printf("i: %d and j: %d\n", numbers[i], numbers[j])
				product = numbers[i] * numbers[j]
				break
			}
		}
	}

	fmt.Printf("Answer: %d\n", product)
}

// SolutionTwo is the second solution for the day
func SolutionTwo() {
	numbers := sidwtrw.SliceOfInts("day01/input.txt")
	product := 0

	sort.Ints(numbers)

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] > 2020 {
					continue
				}
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					fmt.Printf("i: %d and j: %d and k: %d\n", numbers[i], numbers[j], numbers[k])
					product = numbers[i] * numbers[j] * numbers[k]
					break
				}
			}
		}
	}

	fmt.Printf("Answer: %d\n", product)
}
