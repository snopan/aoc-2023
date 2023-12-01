package day1

import (
	"testing"

	"gotest.tools/assert"
)

func Test_SolutionPart1(t *testing.T) {
	input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
	solution := SolutionPart1(input)
	assert.Equal(t, solution, 142)

	input = ""
	solution = SolutionPart1(input)
	assert.Equal(t, solution, 0)
}

func Test_SolutionPart2(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	solution := SolutionPart2(input)
	assert.Equal(t, solution, 281)

	input = ""
	solution = SolutionPart2(input)
	assert.Equal(t, solution, 0)
}

func Test_convertWordNumbers(t *testing.T) {
	input := "9fgsixzkbscvbxdsfive6spjfhzxbzvgbvrthreeoneightn"
	output := convertWordNumbers(input)
	assert.Equal(t, output, "9fg6zkbscvbxds56spjfhzxbzvgbvr318n")
}

func Test_generateWordNumberTree(t *testing.T) {
	tree := generateWordNumberTree()
	two := tree.children["t"].children["w"].children["o"].value
	nine := tree.children["n"].children["i"].children["n"].children["e"].value
	assert.Equal(t, two, 2)
	assert.Equal(t, nine, 9)
}
