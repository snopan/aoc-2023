package day4

import (
	"testing"

	"github.com/snopan/aoc-2023/src/helpers"
	"gotest.tools/assert"
)

func Test_SolutionPart1(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
	`

	solution := SolutionPart1(input)
	assert.Equal(t, solution, 13)
}

func Test_SolutionPart2(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
	Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
	Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
	Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
	Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
	`

	solution := SolutionPart2(input)
	assert.Equal(t, solution, 30)
}

func Test_parseCardNumbers(t *testing.T) {
	input := "1 2 3 4"
	output := helpers.GetNumbersFromText(input)
	assert.Equal(t, output[0], 1)
	assert.Equal(t, output[1], 2)
	assert.Equal(t, output[2], 3)
	assert.Equal(t, output[3], 4)
}
