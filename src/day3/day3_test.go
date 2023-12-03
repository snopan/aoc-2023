package day3

import (
	"testing"

	"gotest.tools/assert"
)

func Test_SolutionPart1(t *testing.T) {
	input := "467..114..\n...*......\n..35*.633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
	solution := SolutionPart1(input)
	assert.Equal(t, solution, 4361)
}

func Test_SolutionPart2(t *testing.T) {
	input := "467..114..\n...*......\n..35*.633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."
	solution := SolutionPart2(input)
	assert.Equal(t, solution, 467835)
}
