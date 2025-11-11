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
