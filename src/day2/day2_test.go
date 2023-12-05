package day2

import (
	"testing"

	"github.com/snopan/aoc-2023/src/helpers"
	"gotest.tools/assert"
)

var input = helpers.PrepareInput(`
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`)

func Test_SolutionPart1(t *testing.T) {
	solution := SolutionPart1(input)
	assert.Equal(t, solution, 8)
}

func Test_SolutionPart2(t *testing.T) {
	solution := SolutionPart2(input)
	assert.Equal(t, solution, 2286)
}

func Test_parseBag(t *testing.T) {
	allMarbles := "12 red, 13 green, 14 blue"
	bag := parseBag(allMarbles)
	assert.Equal(t, bag.red, 12)
	assert.Equal(t, bag.green, 13)
	assert.Equal(t, bag.blue, 14)

	onlyRed := "12 red"
	bag = parseBag(onlyRed)
	assert.Equal(t, bag.red, 12)
	assert.Equal(t, bag.green, 0)
	assert.Equal(t, bag.blue, 0)

	empty := ""
	bag = parseBag(empty)
	assert.Equal(t, bag.red, 0)
	assert.Equal(t, bag.green, 0)
	assert.Equal(t, bag.blue, 0)
}
