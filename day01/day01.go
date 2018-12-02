package day01

// This was done in bash as follows:
// cat input.txt | xargs | sed -e 's/^+//' | bc

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// HandleError is a generic error handler
func HandleError(err error) {
	if err != nil {
		fmt.Print(err)
		os.Exit(2)
	}
}

// SliceOfNumbers returns a slice of numbers from the input file
func SliceOfNumbers() []int {
	file, err := os.Open("day01/input.txt")
	HandleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		number := scanner.Text()
		num, err := strconv.Atoi(number)
		HandleError(err)
		numbers = append(numbers, num)
	}
	return numbers
}

// SolutionOne is the first solution for the day
func SolutionOne() {
	numbers := SliceOfNumbers()
	sum := 0
	for _, num := range numbers {
		sum += num
	}

	fmt.Printf("Answer: %d\n", sum)
}

// SolutionTwo is the second solution for the day
func SolutionTwo() {
	numbers := SliceOfNumbers()

	highestNumber := false
	numberMap := make(map[int]int)
	sum := 0

	for {
		for _, num := range numbers {
			sum += num
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
