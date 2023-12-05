package day5

import (
	"testing"

	"github.com/snopan/aoc-2023/src/helpers"
	"gotest.tools/assert"
)

var input = helpers.PrepareInput(`
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`)

func Test_SolutionPart1(t *testing.T) {
	solution := SolutionPart1(input)
	assert.Equal(t, solution, 35)
}

func Test_SolutionPart2(t *testing.T) {
	solution := SolutionPart2(input)
	assert.Equal(t, solution, 46)
}

func Test_parseMapping(t *testing.T) {
	input := helpers.PrepareInput(`
seed-to-soil map:
50 98 2
52 50 48`)
	output := parseMapping(input)
	assert.Equal(t, output[0].GetDest(98), 50)
	assert.Equal(t, output[0].GetDest(99), 51)
	assert.Equal(t, output[1].GetDest(50), 52)
	assert.Equal(t, output[1].GetDest(53), 55)
}
