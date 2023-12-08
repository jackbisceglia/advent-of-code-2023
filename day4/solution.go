// Day 4
// go run .

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func split_line_by_pipe(line string) (winning_nums []int, num []int) {
	// find all the numbers in the string, convert each to int, copy over, return
	parse_nums_str := func(nums string) []int {
		nums_str_arr := regexp.MustCompile(`\d+`).FindAllString(nums, -1)

		nums_int_arr := []int{}
		for _, num_str := range nums_str_arr {
			num, _ := strconv.Atoi(num_str)
			nums_int_arr = append(nums_int_arr, num)
		}

		return nums_int_arr
	}

	// we first remove the card prefix, then split into 2 by pipe
	nums_only := regexp.MustCompile(`Card.+\d+:`).ReplaceAllString(line, "")
	split := strings.Split(nums_only, "|")

	// then we return int array of each
	return parse_nums_str(split[0]), parse_nums_str(split[1])
}

func count_num_matches(slice []int, num int) int {
	count := 0
	for _, n := range slice {
		if n == num {
			count += 1
		}
	}
	return count
}

func solve_p1(puzzle []string) int {
	sum := 0
	// grab the two sets of nums, count the matches, then add 2^matches-1 to sum
	for _, line := range puzzle {
		winning_nums, nums := split_line_by_pipe(line)

		matches := 0

		for _, winning_num := range winning_nums {
			matches += count_num_matches(nums, winning_num)
		}

		if matches >= 0 {
			sum += int(math.Pow(float64(2), float64(matches-1)))
		}
	}

	return sum
}

func solve_p2(puzzle []string) int {
	sum := 0
	card_idx_to_copies_map := map[int]int{}

	// map the card number -> number of copies found
	//   - since we can only get more copies of future cards, we can do this all in 1 pass
	// then we increment count to the subsequent <matches> cards
	// finally we increment our sum by the amount of copies found thus far + the original copy
	// (we index by idx + 1 since card numbers start at 1)
	for line_idx, line := range puzzle {
		winning_nums, nums := split_line_by_pipe(line)

		matches := 0

		for _, winning_num := range winning_nums {
			matches += count_num_matches(nums, winning_num)
		}

		for i := 1; i <= matches; i++ {
			card_idx_to_copies_map[line_idx+i+1] += card_idx_to_copies_map[line_idx+1] + 1
		}

		sum += card_idx_to_copies_map[line_idx+1] + 1
	}

	return sum
}

func main() {
	fmt.Println("Running Code - Day 4")

	puzzle := parse()

	fmt.Printf("Day 1: %d\n", solve_p1(puzzle))
	fmt.Printf("Day 2: %d\n", solve_p2(puzzle))
}
