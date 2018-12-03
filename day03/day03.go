package day03

import (
	"fmt"

	"github.com/chazdnato/aoc/sidwtrw"
)

// SolutionOne is the first solution for the day
func SolutionOne() {
	claims := sidwtrw.SliceOfStrings("day03/input.txt")
	const dx, dy = 1000, 1000
	fabric := [dx][dy]int{}

	// parse line like '#62 @ 301,425: 27x25' as
	// #(62).@.(301),(425):.(27)x(25) where the groups are
	// entry number, x-coord, y-coord, width, height (entry, x, y, w, h)
	regex := `#(?P<entry>\d+).@.(?P<x>\d+),(?P<y>\d+):.(?P<w>\d+)x(?P<h>\d+)`
	for _, claim := range claims {
		claimMap := sidwtrw.ParseRegex(regex, claim)

		xcoord := sidwtrw.ConvertStrtoInt(claimMap["x"])
		xrange := sidwtrw.ConvertStrtoInt(claimMap["w"])
		ycoord := sidwtrw.ConvertStrtoInt(claimMap["y"])
		yrange := sidwtrw.ConvertStrtoInt(claimMap["h"])
		for x := xcoord; x < xcoord+xrange; x++ {
			for y := ycoord; y < ycoord+yrange; y++ {
				fabric[x][y]++
			}
		}

	}

	total := 0
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			if fabric[x][y] > 1 {
				total++
			}
		}
		// fmt.Printf("%v\n", fabric[x])
	}
	fmt.Printf("Answer: %d\n", total)
}

// SolutionTwo is the second solution for the day
func SolutionTwo() {
	claims := sidwtrw.SliceOfStrings("day03/input.txt")
	const dx, dy = 1000, 1000
	fabric := [dx][dy]int{}

	trackUnclaimed := make(map[int]int)

	// parse line like '#62 @ 301,425: 27x25' as
	// #(62).@.(301),(425):.(27)x(25) where the groups are
	// entry number, x-coord, y-coord, width, height (entry, x, y, w, h)
	regex := `#(?P<entry>\d+).@.(?P<x>\d+),(?P<y>\d+):.(?P<w>\d+)x(?P<h>\d+)`
	for _, claim := range claims {
		claimMap := sidwtrw.ParseRegex(regex, claim)

		xcoord := sidwtrw.ConvertStrtoInt(claimMap["x"])
		xrange := sidwtrw.ConvertStrtoInt(claimMap["w"])
		ycoord := sidwtrw.ConvertStrtoInt(claimMap["y"])
		yrange := sidwtrw.ConvertStrtoInt(claimMap["h"])
		entry := sidwtrw.ConvertStrtoInt(claimMap["entry"])
		for x := xcoord; x < xcoord+xrange; x++ {
			for y := ycoord; y < ycoord+yrange; y++ {
				// claimed entry
				if fabric[x][y] != 0 {
					// invalidate previous claim and current entry
					trackUnclaimed[fabric[x][y]] = -1
					trackUnclaimed[entry] = -1
				} else {
					// only claim if it's never been falsified
					if trackUnclaimed[entry] != -1 {
						trackUnclaimed[entry] = 1
					}
				}
				// mark your territory
				fabric[x][y] = entry
			}
		}
	}
	for entry, truthiness := range trackUnclaimed {
		if truthiness == 1 {
			fmt.Printf("Answer: %d\n", entry)
		}
	}
}
