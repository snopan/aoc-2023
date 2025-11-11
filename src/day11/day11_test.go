package day11

import (
	"testing"

	"github.com/snopan/aoc-2023/src/helpers"
	"gotest.tools/assert"
)

var input = helpers.PrepareInput(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`)

func Test_SolutionPart1(t *testing.T) {
	solution := SolutionPart1(input)
	assert.Equal(t, solution, 374)
}

func Test_SolutionPart2(t *testing.T) {
	assert.Equal(t, commonSolution(input, 10), 1030)
	assert.Equal(t, commonSolution(input, 100), 8410)
}
