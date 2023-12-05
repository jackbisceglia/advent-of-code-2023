// Day 2
// go run .

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

// split line into ("Game <x>: ", and the views of cubes string)
func extract_game_ident(line string) (string, string) {
	game_ident_regex, _ := regexp.Compile(`\bGame \d(\d)?(\d)?: `)

	game_id := game_ident_regex.FindString(line)
	game_info := game_ident_regex.ReplaceAllString(line, "")

	return game_id, game_info
}

// extract rgb given a single cube view (line delimited by ;)
func extract_cubes(cube_view string) (int, int, int) {
	gen_regex := regexp.MustCompile(`\D`)
	r_regex, _ := regexp.Compile(`\d+ red`)
	g_regex, _ := regexp.Compile(`\d+ green`)
	b_regex, _ := regexp.Compile(`\d+ blue`)

	r, _ := strconv.Atoi(gen_regex.ReplaceAllString(r_regex.FindString(cube_view), ""))
	g, _ := strconv.Atoi(gen_regex.ReplaceAllString(g_regex.FindString(cube_view), ""))
	b, _ := strconv.Atoi(gen_regex.ReplaceAllString(b_regex.FindString(cube_view), ""))

	return r, g, b
}

type Bag struct {
	r int
	g int
	b int
}

func solve_p1(puzzle []string) int {
	bagDetails := Bag{
		r: 12,
		g: 13,
		b: 14,
	}

	// code here
	id_sum := 0
	// loop over, add game_id every time, and subtract if we find a cube count too big
	// easier to just subtract when necessary rather than keeping track of sum flag to see overflow and then add
	for i, line := range puzzle {
		_, game_info := extract_game_ident(line)
		id_sum += i + 1

		views := strings.Split(game_info, ";")
		for _, cube_view := range views {
			r, g, b := extract_cubes(cube_view)

			if r > bagDetails.r || g > bagDetails.g || b > bagDetails.b {
				id_sum -= i + 1
				break
			}
		}

	}

	return id_sum
}

func (b *Bag) calc_power() int {
	return b.r * b.g * b.b
}

func solve_p2(puzzle []string) int {
	product_sum := 0
	// for each game we just keep track of the max cubes we've seen, and tally the product at the end
	for _, line := range puzzle {
		_, game_info := extract_game_ident(line)

		max_cubes := Bag{
			r: 0,
			g: 0,
			b: 0,
		}

		views := strings.Split(game_info, ";")
		for _, cube_view := range views {
			r, g, b := extract_cubes(cube_view)
			if r > max_cubes.r {
				max_cubes.r = r
			}
			if g > max_cubes.g {
				max_cubes.g = g
			}
			if b > max_cubes.b {
				max_cubes.b = b
			}
		}

		product_sum += max_cubes.calc_power()
	}

	return product_sum
}

func main() {
	fmt.Println("Running Code - Day 2")

	puzzle := parse()

	fmt.Printf("Day 1: %d\n", solve_p1(puzzle))
	fmt.Printf("Day 2: %d\n", solve_p2(puzzle))
}
