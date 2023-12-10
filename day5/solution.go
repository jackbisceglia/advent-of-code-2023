// Day 5
// go run .

package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
)

var PUZZLE_PATH = "puzzle.txt"
var SAMPLE_PATH = "sample.txt"

func parse() []string {
	// file, err := os.ReadFile(SAMPLE_PATH)
	file, err := os.ReadFile(PUZZLE_PATH)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(strings.Trim(string(file), "\n"), "\n\n")
}

func seeds_to_bigint(seeds []string) []big.Int {
	seeds_as_int := []big.Int{}
	for _, seed_string := range seeds {
		seed, _ := new(big.Int).SetString(seed_string, 10) // strconv.Atoi(seed_string)
		seeds_as_int = append(seeds_as_int, *seed)
	}
	return seeds_as_int
}

// data structure(s) to make defining the ranges and converting a bit easier
// not really needed, but helps avoided indexing deep into 2len arrays without names
// instead can do convert_map.Ranges[0].Src.min, etc
type _Range struct {
	min big.Int
	max big.Int
}

type ConvertRange struct {
	Src  _Range
	Dest _Range
}

type ConvertMap struct {
	Ranges []ConvertRange
}

func (c ConvertMap) convert_reverse(input big.Int) big.Int {
	in_range := func(cr ConvertRange, input big.Int) bool {
		return input.Cmp(&cr.Dest.min) >= 0 && input.Cmp(&cr.Dest.max) <= 0
	}

	// run through our ranges and if the input is in bounds, convert it. else, return input
	for _, r := range c.Ranges {
		if in_range(r, input) {
			offset := *big.NewInt(0).Sub(&input, &r.Dest.min)
			return *big.NewInt(0).Add(&r.Src.min, &offset)
		}
	}

	return input
}

func (c ConvertMap) convert(input big.Int) big.Int {
	in_range := func(cr ConvertRange, input big.Int) bool {
		return input.Cmp(&cr.Src.min) >= 0 && input.Cmp(&cr.Src.max) <= 0
	}

	// run through our ranges and if the input is in bounds, convert it. else, return input
	for _, r := range c.Ranges {
		if in_range(r, input) {
			offset := *big.NewInt(0).Sub(&input, &r.Src.min)
			return *big.NewInt(0).Add(&r.Dest.min, &offset)
		}
	}

	return input
}

func create_convert_map(input []string) ConvertMap {
	cm := ConvertMap{
		Ranges: []ConvertRange{},
	}
	// go through each range as string, and construct a ConvertMap struct
	for _, str_range := range input {
		range_arr := strings.Split(str_range, " ")

		src_start_bigint, _ := new(big.Int).SetString(range_arr[1], 10)
		dest_start_bigint, _ := new(big.Int).SetString(range_arr[0], 10)
		length_bigint, _ := new(big.Int).SetString(range_arr[2], 10)

		cr := ConvertRange{
			Src: _Range{
				min: *src_start_bigint,
				max: *big.NewInt(0).Add(src_start_bigint, length_bigint),
			},
			Dest: _Range{
				min: *dest_start_bigint,
				max: *big.NewInt(0).Add(dest_start_bigint, length_bigint),
			},
		}

		cm.Ranges = append(cm.Ranges, cr)
	}

	return cm
}

func parse_puzzle_as_convert_maps(puzzle []string) []ConvertMap {
	maps := []ConvertMap{}
	puzzle = puzzle[1:] // don't need seeds line

	// takes in our array of range sections, and creates a ConvertMap for each
	for _, ranges := range puzzle {
		maps = append(maps, create_convert_map(strings.Split(ranges, "\n")[1:]))
	}

	return maps
}

func solve_p1(puzzle []string) big.Int {
	seeds := seeds_to_bigint(strings.Split(puzzle[0], " ")[1:])

	maps := parse_puzzle_as_convert_maps(puzzle)
	min := big.NewInt(-1)

	// for each seed, run it through each ConverMap and get the final location
	// only track the min
	for _, seed := range seeds {
		conversion := seed
		for _, cm := range maps {
			conversion = cm.convert(conversion)
		}

		if min.Cmp(big.NewInt(-1)) == 0 || conversion.Cmp(min) == -1 {
			min = &conversion
		}
	}

	return *min
}

func get_seed_ranges(seeds []big.Int) []_Range {
	seed_ranges := []_Range{}

	for i := 0; i < len(seeds); i += 2 {
		seed_ranges = append(seed_ranges, _Range{
			min: seeds[i],
			max: *big.NewInt(0).Add(&seeds[i], &seeds[i+1]),
		})
	}
	return seed_ranges
}

func reverse_maps(maps []ConvertMap) []ConvertMap {
	reversed_maps := []ConvertMap{}

	for i := len(maps) - 1; i >= 0; i-- {
		reversed_maps = append(reversed_maps, maps[i])
	}

	return reversed_maps
}

func solve_p2(puzzle []string) big.Int {
	seeds_arr := get_seed_ranges(seeds_to_bigint(strings.Split(puzzle[0], " ")[1:]))
	maps := reverse_maps(parse_puzzle_as_convert_maps(puzzle))

	// start at 1, and go up until we find a location that has a corresponding seed
	// can probably trim down some of the input space, but not sure how
	loc := big.NewInt(1)
	for {
		convert := *loc
		for _, cm := range maps {
			convert = cm.convert_reverse(convert)
		}

		// check if this location is in any of the seeds
		for _, seed := range seeds_arr {
			if convert.Cmp(&seed.min) >= 0 && convert.Cmp(&seed.max) < 0 {
				return *loc
			}
		}

		loc = big.NewInt(0).Add(loc, big.NewInt(1))
	}
}

func main() {
	fmt.Println("Running Code - Day 5")

	puzzle := parse()

	p1_sol := solve_p1(puzzle)
	p2_sol := solve_p2(puzzle)

	fmt.Printf("Day 1: %s\n", p1_sol.String())
	fmt.Printf("Day 2: %s\n", p2_sol.String())
}
