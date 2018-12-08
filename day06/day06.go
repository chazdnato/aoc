package day06

import (
	"fmt"

	"github.com/chazdnato/aoc/sidwtrw"
)

// SolutionOne is the first solution for the day
func SolutionOne() {
	points := sidwtrw.SliceOfStrings("day06/input.txt")
	// total "space" size
	const dx, dy = 400, 400
	areas := make(map[string]int)
	infinite := make(map[string]bool)

	regex := `(?P<xcoord>\d+), (?P<ycoord>\d+)`

	// pre-seed the areas and if something has infinite area
	// going to key the point as a string 'cause it works
	for _, point := range points {
		areas[point] = 0
		infinite[point] = false
	}

	// naively go through each coordinate and calculate
	// the distance from that coordinate to each point
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			minPoint := "0, 0"     // nonce
			minDist := dx + dy + 1 // set to max
			for _, point := range points {
				pointMap := sidwtrw.ParseRegex(regex, point)
				xcoord := sidwtrw.StrToInt(pointMap["xcoord"])
				ycoord := sidwtrw.StrToInt(pointMap["ycoord"])

				// Manhattan distance
				dist := sidwtrw.Abs(xcoord-x) + sidwtrw.Abs(ycoord-y)
				if dist < minDist {
					minDist = dist
					minPoint = point
				} else if dist == minDist { // overlap
					minPoint = "-1, -1"
				}
			}
			// bounds check
			if x == dx-1 || x == 0 || y == dx-1 || y == 0 {
				infinite[minPoint] = true
			}
			// count area for point w/o conflict, or dump in "-1, -1"
			areas[minPoint]++
		}
	}

	maxArea := 0
	for point, count := range areas {
		if count > maxArea && !infinite[point] {
			maxArea = count
		}
	}

	fmt.Printf("Answer: %v\n", maxArea)
}

// SolutionTwo is the second solution for the day
func SolutionTwo() {
	coords := sidwtrw.SliceOfStrings("day06/input_sample1.txt")

	fmt.Printf("Answer: %v\n", coords)
}
