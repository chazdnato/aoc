package day02

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/chazdnato/aoc/sidwtrw"
)

var COLOURS = map[string]int{
	"blue":  14,
	"green": 13,
	"red":   12,
}

func parseCube(round string, colour string) int {
	re := regexp.MustCompile(fmt.Sprintf(`(\d+) %s`, colour))
	matches := re.FindStringSubmatch(round)
	digit := 0
	if len(matches) > 1 {
		digitStr := matches[1]
		digit, _ = strconv.Atoi(digitStr)
	}
	return digit
}

func extractGamedata(game string) (gameNum int, rounds string) {
	gameSplit := strings.Split(game, ":")
	gameString := strings.Split(gameSplit[0], " ")[1]
	gameNum, err := strconv.Atoi(gameString)
	sidwtrw.HandleError("converting gamenum", err)
	rounds = gameSplit[1]
	return gameNum, rounds
}

// SolutionOne is the first solution for the day
func SolutionOne() {
	lines := sidwtrw.SliceOfStrings("day02/input.txt")
	total := 0

	for _, line := range lines {
		gameNum, rounds := extractGamedata(line)
		passed := true

		for _, round := range strings.Split(rounds, ";") {
			// fmt.Printf("%d - %s\n", game_num, round)
			for colour, maxMarbles := range COLOURS {
				digit := parseCube(round, colour)
				if digit > maxMarbles {
					passed = false
					// fmt.Printf("\t\tThere are too many %s cubes\n", colour)
				}
				// fmt.Printf("There are %d %s\n", digit, colour)
			}
		}

		if passed {
			total += gameNum
			// fmt.Printf("==== ROUND %d PASSED ===\n", game_num)
		} else {
			// fmt.Printf("==== ROUND %d FAILED ===\n", game_num)
		}

	}

	fmt.Printf("%d\n", total)
}

// SolutionOne is the first solution for the day
func SolutionTwo() {
	lines := sidwtrw.SliceOfStrings("day02/input.txt")
	total := 0

	for _, game := range lines {
		_, rounds := extractGamedata(game)
		gameProduct := 1

		// track the highest number of coloured cubes
		maxCubes := map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}
		// fmt.Printf("\nGame Number: %d\n", gameNum)
		for _, round := range strings.Split(rounds, ";") {
			// fmt.Printf("\nRound: %s\n", round)

			// for each colour type, look at how many of that kind of cube
			// is in this round, and store the highest
			for colour, numCubes := range maxCubes {
				digit := parseCube(round, colour)
				if digit > numCubes {
					maxCubes[colour] = digit
				}
				// fmt.Printf("\t%v\t", maxCubes)
			}
		}

		// now get product of all cube maximii
		for _, max := range maxCubes {
			gameProduct *= max
		}

		// add to total
		total += gameProduct
	}

	fmt.Printf("%d\n", total)
}
