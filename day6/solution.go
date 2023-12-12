// Day 6
// go run .

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"regexp"
	"strconv"
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

func quadratic_formula(a, b, c float64) (float64, float64) {
	d := float64(b*b - 4*a*c)
	var root1, root2 float64

	if d < 0 {
		root1, root2 = float64(-1), float64(-1)
	} else if d == 0 {
		root1, root2 = (-b+math.Sqrt(d))/(2*a), (-b+math.Sqrt(d))/(2*a)
	} else {
		root1, root2 = (-b+math.Sqrt(d))/(2*a), (-b-math.Sqrt(d))/(2*a)
	}

	return root1, root2
}

func solve_p1(puzzle []string) int {
	// grab all the numbers, iterate over them in parallel since their lengths are equal
	// at each step, calculate roots of the quadratic, and these are the bound of winning strategies
	// do some rounding and return the difference
	times := regexp.MustCompile(`\d+`).FindAllString(puzzle[0], -1)
	distances := regexp.MustCompile(`\d+`).FindAllString(puzzle[1], -1)
	total := 1

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])

		// in order to handle the strictly > condition, we round to get 1 out of range
		// once we have those values, we add/subtract 1 respectively to move values in range
		root1, root2 := quadratic_formula(float64(1), float64(-time), float64(distance))
		min_root, max_root := math.Min(root1, root2), math.Max(root1, root2)
		low_rounded, high_rounded := math.Floor(min_root)+1, math.Ceil(max_root)-1

		total *= (int(high_rounded) - int(low_rounded) + 1)
	}

	return total
}

func condense_nums(str string) string {
	is_digit := func(c rune) bool { return c >= '0' && c <= '9' }
	condense := ""

	for _, c := range str {
		if is_digit(c) {
			condense += string(c)
		}
	}

	return condense
}

func solve_p2(puzzle []string) int {
	// same core logic as p1, but we need to transform the input differently
	// we first condense the numbers by removing non-digits, then cast to big int
	// to handle numbers primitive int range, then we cast to float64 to use the quad helpers
	time_bigint, _ := big.NewInt(0).SetString(condense_nums(puzzle[0]), 10)
	distance_bigint, _ := big.NewInt(0).SetString(condense_nums(puzzle[1]), 10)

	time, _ := time_bigint.Float64()
	distance, _ := distance_bigint.Float64()

	root1, root2 := quadratic_formula(float64(1), -time, float64(distance))
	min_root, max_root := math.Min(root1, root2), math.Max(root1, root2)
	low_rounded, high_rounded := math.Floor(min_root)+1, math.Ceil(max_root)-1

	return int(high_rounded - low_rounded + 1)
}

func main() {
	fmt.Println("Running Code - Day 6")

	puzzle := parse()

	fmt.Printf("Day 1: %d\n", solve_p1(puzzle))
	fmt.Printf("Day 2: %d\n", solve_p2(puzzle))
}
