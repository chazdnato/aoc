// Package sidwtrw - Shit I Don't Want To Re-Write
package sidwtrw

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Abs will give absolute function for integers
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// StrToInt will convert a string to an int
func StrToInt(str string) int {
	num, err := strconv.Atoi(str)
	HandleError("Converting string to integer", err)
	return num
}

// HandleError is a generic error handler
func HandleError(location string, err error) {
	if err != nil {
		fmt.Printf("ERR: %s\t %v\n", location, err)
		os.Exit(2)
	}
}

// ParseRegex will use a regex to split a string into a map of entries
// stolen blatantly from https://stackoverflow.com/a/39635221
func ParseRegex(regEx, toParse string) (parsedMap map[string]string) {

	var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(toParse)

	parsedMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i == 0 {
			name = "_ALL_"
		}
		parsedMap[name] = match[i]
	}
	return
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
		numbers = append(numbers, StrToInt(number))
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
