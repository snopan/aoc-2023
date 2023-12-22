package day10

import (
	"testing"

	"github.com/snopan/aoc-2023/src/helpers"
	"gotest.tools/assert"
)

var input = helpers.PrepareInput(`
7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ
`)

var input2 = helpers.PrepareInput(`
...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........
`)

func Test_SolutionPart1(t *testing.T) {
	solution := SolutionPart1(input)
	assert.Equal(t, solution, 8)
}

func Test_SolutionPart2(t *testing.T) {
	solution := SolutionPart2(input2)
	assert.Equal(t, solution, 4)
}
