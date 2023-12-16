package day9

import (
	"testing"

	"github.com/snopan/aoc-2023/src/helpers"
	"gotest.tools/assert"
)

var input = helpers.PrepareInput(`
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`)

func Test_SolutionPart1(t *testing.T) {
	solution := SolutionPart1(input)
	assert.Equal(t, solution, 114)
}

func Test_SolutionPart2(t *testing.T) {
	solution := SolutionPart2(input)
	assert.Equal(t, solution, 2)
}
