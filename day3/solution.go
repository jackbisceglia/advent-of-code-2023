// Day 3
// go run .

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var PUZZLE_PATH = "puzzle.txt"
var SAMPLE_PATH = "sample.txt"

func parse() []string {
	// read
	// file, err := os.Open(SAMPLE_PATH)
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

var directions = [][]int{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

func is_symbol(char rune) bool {
	return char != '.' && !unicode.IsDigit(char)
}

func is_digit(char rune) bool {
	return unicode.IsDigit(char)
}

func get_first_number(sub_line string) string {
	return regexp.MustCompile(`\d+`).FindString(sub_line)
}

func position_is_in_bounds(matrix []string, x int, y int) bool {
	return x >= 0 && x < len(matrix[0]) && y >= 0 && y < len(matrix)
}

func number_has_adjacent_symbol(matrix []string, x int, y int, number_as_string string) bool {
	digit_has_adjacent_symbol := func(x int, y int) bool {
		for _, direction := range directions {

			new_x := x + direction[0]
			new_y := y + direction[1]

			if position_is_in_bounds(matrix, new_x, new_y) && is_symbol(rune(matrix[new_y][new_x])) {
				return true
			}
		}
		return false
	}

	for offset := range number_as_string {
		if digit_has_adjacent_symbol(x+offset, y) {
			return true
		}
	}

	return false
}

func solve_p1(puzzle []string) int {
	sum := 0
	for y, line := range puzzle {
		x := 0
		for x < len(line) {
			char := line[x]
			current_number_str := get_first_number(line[x:])
			current_number, _ := strconv.Atoi(current_number_str)

			if is_digit(rune(char)) {
				if number_has_adjacent_symbol(puzzle, x, y, current_number_str) {
					sum += current_number
				}

				x += len(current_number_str)
			} else {
				x += 1
			}
		}
	}

	return sum
}

// takes in a gear symbol and returns the ratio or false
func get_gear_ratio_by_symbol(matrix []string, x int, y int) (int, bool) {
	scan_left_and_grab_number := func(x int, y int) (num int, left_idx int) {
		i := x
		for i > 0 && is_digit(rune(matrix[y][i])) {
			i -= 1
		}
		full_number, _ := strconv.Atoi(get_first_number(matrix[y][i:]))

		return full_number, i
	}

	is_duplicate := func(adj_list [][]int, num int, idx int) bool {
		for _, adj := range adj_list {
			if adj[0] == num && adj[1] == idx {
				return true
			}
		}
		return false
	}

	// track adj numbers (and their left most indices to make sure we don't have dupes)
	adjacent_nums := [][]int{}
	for _, direction := range directions {
		new_x := x + direction[0]
		new_y := y + direction[1]

		if !position_is_in_bounds(matrix, new_x, new_y) || !is_digit(rune(matrix[new_y][new_x])) {
			continue
		}

		// grab num
		full_number, left_idx := scan_left_and_grab_number(new_x, new_y)

		if !is_duplicate(adjacent_nums, full_number, left_idx) {
			adjacent_nums = append(adjacent_nums, []int{full_number, left_idx})
		}
	}

	if len(adjacent_nums) != 2 {
		return -1, false
	}

	return adjacent_nums[0][0] * adjacent_nums[1][0], true
}

func solve_p2(puzzle []string) int {
	sum := 0
	// scan through puzzle for potential gears (asterisks), then check if they're gears and sum them
	for y, line := range puzzle {
		for x := range line {
			if rune(puzzle[y][x]) != '*' {
				continue
			}

			current_gear_ratio, is_gear := get_gear_ratio_by_symbol(puzzle, x, y)

			if is_gear {
				sum += current_gear_ratio
			}
		}
	}

	return sum
}

func main() {
	fmt.Println("Running Code - Day 3")

	puzzle := parse()

	fmt.Printf("Day 1: %d\n", solve_p1(puzzle))
	fmt.Printf("Day 2: %d\n", solve_p2(puzzle))
}
