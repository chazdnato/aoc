package day07

import (
	"fmt"
	"sort"

	"github.com/chazdnato/aoc/sidwtrw"
)

// remove will take the last element and replace the element
// you desire to remove, and return the truncated slice
func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// indexOf returns the index of a string in a slice of strings
func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found
}

// sortMapKeys will take a map[string][]string
// and return sorted keys in []string
func sortMapKeys(m map[string][]string) []string {
	o := make([]string, 0, len(m))
	for k := range m {
		o = append(o, k)
	}
	sort.Strings(o) //sort by key
	return o
}

// SolutionOne is the first solution for the day
func SolutionOne() {
	stepMap := make(map[string][]string)
	steps := sidwtrw.SliceOfStrings("day07/input.txt")

	// populate steps
	for _, step := range steps {
		prec := string(step[5])
		dec := string(step[36])

		// initialize
		if _, ok := stepMap[prec]; !ok {
			stepMap[prec] = []string{}
		}
		if _, ok := stepMap[dec]; !ok {
			stepMap[dec] = []string{}
		}
		// populate
		stepMap[dec] = append(stepMap[dec], prec)
	}

	// build string
	stepOrder := ""
	haveSteps := true
	for haveSteps {
		// sort the keys
		orderedSteps := sortMapKeys(stepMap)
		for _, step := range orderedSteps {
			// find the first empty key
			if len(stepMap[step]) == 0 {
				// append key to stepOrder
				stepOrder += step
				// for all keys, remove empty key from arrays
				for k, v := range stepMap {
					// skip current step, already empty
					if k == step {
						continue
					}
					// remove item from key map if it exists
					if i := indexOf(step, v); i >= 0 {
						stepMap[k] = remove(stepMap[k], i)
					}
				}
				// remove key from stepMap then start over
				delete(stepMap, step)
				break
			}
		}
		if len(stepMap) == 0 {
			haveSteps = false
		}
	}

	fmt.Printf("Answer: %v\n", stepOrder)
}

// SolutionTwo is the second solution for the day
func SolutionTwo() {
	steps := sidwtrw.SliceOfStrings("day07/input.txt")

	fmt.Printf("Answer: %v\n", steps)
}
