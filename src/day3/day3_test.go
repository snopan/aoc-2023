package day3

import (
	"testing"

	"github.com/snopan/aoc-2023/src/helpers"
	"gotest.tools/assert"
)

var input = helpers.PrepareInput(`
467..114..
...*......
..35*.633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

func Test_SolutionPart1(t *testing.T) {
	solution := SolutionPart1(input)
	assert.Equal(t, solution, 4361)
}

func Test_SolutionPart2(t *testing.T) {
	solution := SolutionPart2(input)
	assert.Equal(t, solution, 467835)
}
