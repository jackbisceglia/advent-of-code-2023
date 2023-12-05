// Day 1
// go run .

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var PUZZLE_PATH = "puzzle.txt"
var SAMPLE_PATH = "sample.txt"

func parse() []string {
	// read
	file, err := os.Open(PUZZLE_PATH)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func solve_p1(puzzle []string) int {
	// code here
	// general idea: we can keep a pointer on both sides of each line and just move until we find a match
	sum := 0

	for _, line := range puzzle {
		left_i := 0
		right_i := len(line) - 1

		for left_i < right_i {
			// move either ptr inwards if it's not on a digit
			_, err_left := strconv.Atoi(string(line[left_i]))
			if err_left != nil {
				left_i += 1
			}

			_, err_right := strconv.Atoi(string(line[right_i]))
			if err_right != nil {
				right_i -= 1
			}

			// if both are on digits, then we break
			if err_left == nil && err_right == nil {
				break
			}
		}

		value_from_line := fmt.Sprintf("%c%c", line[left_i], line[right_i])
		value_as_string, err := strconv.Atoi(value_from_line)

		if err != nil {
			log.Fatal(err)
		}

		sum += value_as_string

	}

	return sum
}

// reverse string util
func reverse_string(s string) string {
	new_str := []rune{}

	for _, char := range s {
		new_str = append([]rune{char}, new_str...)
	}

	return string(new_str)
}

// check if the substring starts with the spelled version of a digit
// we can pass in regalar and reversed versions of each param to search from both sides
func has_spelled_digit_prefix(substring string, digits_spelled []string) (int, error) {

	for i, digit_spelled := range digits_spelled {
		if strings.HasPrefix(substring, digit_spelled) {
			return i + 1, nil
		}
	}

	return -1, fmt.Errorf("no digit found")

}

func get_first_digit_from_str(line string, reversed bool) int {
	digits_spelled := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	digits_spelled_reverse := []string{
		reverse_string("one"),
		reverse_string("two"),
		reverse_string("three"),
		reverse_string("four"),
		reverse_string("five"),
		reverse_string("six"),
		reverse_string("seven"),
		reverse_string("eight"),
		reverse_string("nine"),
	}

	// init the line to parse and digits slice depending on reversed flag
	var line_to_parse string
	var digits_slice []string

	if reversed {
		line_to_parse = reverse_string(line)
		digits_slice = digits_spelled_reverse
	} else {
		line_to_parse = line
		digits_slice = digits_spelled
	}

	// for each char, if it's a digit return, otherwise, check if it's a spelled digit and return
	for i, char := range line_to_parse {
		if unicode.IsNumber(char) {
			digit, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatal(err)
			}
			return digit
		}

		// get the spelled digit, pass in the forwards or backwards string and digits depending on reversed flag
		digit_spelled, err := has_spelled_digit_prefix(line_to_parse[i:], digits_slice)
		if err == nil {
			return digit_spelled
		}

	}
	return -1
}

func solve_p2(puzzle []string) int {
	// code here
	// retrieve

	sum := 0

	for _, line := range puzzle {
		value_as_string, err := strconv.Atoi(fmt.Sprintf("%d%d", get_first_digit_from_str(line, false), get_first_digit_from_str(line, true)))

		if err != nil {
			log.Fatal(err)
		}

		sum += value_as_string
	}

	return sum
}

func main() {
	fmt.Println("Running Code - Day 1")

	puzzle := parse()

	fmt.Printf("Day 1: %d\n", solve_p1(puzzle))
	fmt.Printf("Day 2: %d\n", solve_p2(puzzle))
}
