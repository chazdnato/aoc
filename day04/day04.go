package day04

import (
	"fmt"
	"sort"

	"github.com/chazdnato/aoc/sidwtrw"
)

type guard struct {
	id      int
	total   int
	minutes [60]int
}

// SolutionOne is the first solution for the day
func SolutionOne() {
	shifts := sidwtrw.SliceOfStrings("day04/input.txt")
	sort.Strings(shifts)
	// key is guard number, value is guard struct
	guards := make(map[int]*guard)

	regex := `.+:(?P<min>\d+). (?P<action>\w+) #?(?P<meta>\S+).+`

	var c *guard // current guard
	startTime, endTime := 0, 0
	for _, shift := range shifts {
		// snag line
		guardMap := sidwtrw.ParseRegex(regex, shift)

		// if guard starts shift, begin read ahead for shift metadata
		if guardMap["action"] == "Guard" {
			guardID := sidwtrw.StrToInt(guardMap["meta"])
			if _, ok := guards[guardID]; !ok {
				guards[guardID] = &guard{id: guardID}
			}
			c = guards[guardID]
		}

		// if the guard falls asleep, capture the nap start
		if guardMap["action"] == "falls" {
			startTime = sidwtrw.StrToInt(guardMap["min"])
		}

		// when the guard wakes, count total sleep time and which minutes asleep
		if guardMap["action"] == "wakes" {
			endTime = sidwtrw.StrToInt(guardMap["min"])
			for x := startTime; x < endTime; x++ {
				c.total++
				c.minutes[x]++
			}
		}
	}

	// find out which guard slept the most
	sg := &guard{total: 0} // sleepiest guard
	for _, guard := range guards {
		if guard.total > sg.total {
			sg = guard
		}
	}

	// find out which minute the guard slept the most
	maxMin := -1
	maxCount := 0
	for i, count := range sg.minutes {
		if count > maxCount {
			maxCount = count
			maxMin = i
		}
	}

	fmt.Printf("Answer: %d\n", sg.id*maxMin)
}

// SolutionTwo is the second solution for the day
func SolutionTwo() {
	fmt.Printf("Answer: <nil>\n")
}
