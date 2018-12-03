// Package sidwtrw - Shit I Don't Want To Re-Write
package sidwtrw

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// HandleError is a generic error handler
func HandleError(location string, err error) {
	if err != nil {
		fmt.Printf("ERR: %s\t %v\n", location, err)
		os.Exit(2)
	}
}

// SliceOfInts returns a slice of ints from the input file
func SliceOfInts(filePath string) []int {
	file, err := os.Open(filePath)
	HandleError("Opening file", err)
	defer file.Close()

	s := bufio.NewScanner(file)

	var numbers []int

	for s.Scan() {
		number := s.Text()
		num, err := strconv.Atoi(number)
		HandleError("Converting string to integer", err)
		numbers = append(numbers, num)
	}
	return numbers
}

// SliceOfStrings returns a slice of strings from the input file
func SliceOfStrings(filePath string) []string {
	file, err := os.Open(filePath)
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
