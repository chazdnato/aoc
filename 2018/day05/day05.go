package day05

import (
	"fmt"
	"math"

	"github.com/chazdnato/aoc/sidwtrw"
)

func reactPolymer(polymer string) string {
	newPolymer := ""
	changed := true
	for changed {
		for i, c := range polymer {
			// check the next character if we're not at the end of the string
			if i+1 < len(polymer) {
				n := polymer[i+1] // next char

				abs := math.Abs(float64(n) - float64(c))
				// characters are "same"
				if int(abs) == 32 {
					newPolymer = polymer[:i] + polymer[i+2:]
					polymer = newPolymer
					changed = true
					break
				}
			}
			changed = false
		}
	}
	return newPolymer
}

// SolutionOne is the first solution for the day
func SolutionOne() int {
	// just want the single line, aka element 0
	polymer := sidwtrw.SliceOfStrings("day05/input.txt")[0]

	newPolymer := reactPolymer(polymer)
	fmt.Printf("Answer: %v\n", len(newPolymer))
	return len(newPolymer)
}

// SolutionTwo is the second solution for the day
func SolutionTwo() int {
	// just want the single line, aka element 0
	polymer := sidwtrw.SliceOfStrings("day05/input.txt")[0]

	shortestPolymerLen := 50000 // length of full polymer

	// for every letter in the alphabet ...
	for l := 65; l <= 90; l++ {
		// remove all instances from the polymer
		newPolymer := ""
		for _, c := range polymer {
			if int(c) == l || int(c) == l+32 {
				continue
			}
			newPolymer += string(c)
		}

		// see the polymer length
		shortPolymer := reactPolymer(newPolymer)
		if len(shortPolymer) < shortestPolymerLen {
			shortestPolymerLen = len(shortPolymer)
		}

	}
	// report the shortest length
	fmt.Printf("Answer: %d\n", shortestPolymerLen)
	return shortestPolymerLen
}
