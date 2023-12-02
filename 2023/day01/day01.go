package day01

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/chazdnato/aoc/sidwtrw"
)

// SolutionOne is the first solution for the day
func SolutionOne() {
	lines := sidwtrw.SliceOfStrings("day01/input.txt")

	total := 0
	for _, value := range lines {
		var digits []int
		for _, char := range value {
			if unicode.IsDigit(char) {
				digit := int(char - '0')
				digits = append(digits, digit)
			}
			// fmt.Printf("%c - ", char)
		}

		// add the first digit as 10s
		// the second digit is the "last" digit in the slice, even if it's length 1
		total += digits[0] * 10
		total += digits[len(digits)-1]
	}

	fmt.Printf("total: %d\n", total)
}

func SolutionTwo() {
	lines := sidwtrw.SliceOfStrings("day01/input.txt")

	total := 0

	// got this solution from someone else because mine broke
	var part2Regexp = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	var part2Regexp2 = regexp.MustCompile(".*(" + part2Regexp.String() + ")")
	var part2Vals = map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, value := range lines {

		first2, _ := part2Vals[part2Regexp.FindString(value)]
		last2, _ := part2Vals[part2Regexp2.FindStringSubmatch(value)[1]]
		v2 := 10*first2 + last2
		total += v2
	}
	// this solution didn't work, because I was replacing words to digits in
	// numerical order, not in the order they appear in the string. As such, a
	// string like "eightwothree" was ending up as "eigh23" instead of "8wo3"
	// for _, value := range lines {
	// 	// first convert strings of digits to actual digits ... then same alg
	// 	// result := ""
	// 	for k, v := range string_to_digit {
	// 		re := regexp.MustCompile(k)

	// 		// fmt.Printf("Attempting to replace %s with %s within %s\n", k, v, value)
	// 		value = re.ReplaceAllString(value, v)
	// 	}

	// 	fmt.Printf("%s\n", value)
	// 	var digits []int
	// 	for _, char := range value {
	// 		if unicode.IsDigit(char) {
	// 			digit := int(char - '0')
	// 			digits = append(digits, digit)
	// 		}
	// 		// fmt.Printf("%c - ", char)
	// 	}

	// 	// add the first digit as 10s
	// 	// the second digit is the "last" digit in the slice, even if it's length 1
	// 	total += digits[0] * 10
	// 	total += digits[len(digits)-1]
	// }

	fmt.Printf("total: %d\n", total)
}

// used in defunct/flawed day 2
// var string_to_digit = map[string]string{
// 	"one":   "1",
// 	"two":   "2",
// 	"three": "3",
// 	"four":  "4",
// 	"five":  "5",
// 	"six":   "6",
// 	"seven": "7",
// 	"eight": "8",
// 	"nine":  "9",
// }
