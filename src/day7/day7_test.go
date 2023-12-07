package day7

import (
	"testing"

	"github.com/snopan/aoc-2023/src/helpers"
	"gotest.tools/assert"
)

var input = helpers.PrepareInput(`
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`)

func Test_SolutionPart1(t *testing.T) {
	solution := SolutionPart1(input)
	assert.Equal(t, solution, 6440)
}

func Test_SolutionPart2(t *testing.T) {
	solution := SolutionPart2(input)
	assert.Equal(t, solution, 5905)
}

func Test_GetHandType(t *testing.T) {
	assert.Equal(t, GetHandType("AAAAA", false), FiveOfAKind)
	assert.Equal(t, GetHandType("AA8AA", false), FourOfAKind)
	assert.Equal(t, GetHandType("23332", false), FullHouse)
	assert.Equal(t, GetHandType("TTT98", false), ThreeOfAKind)
	assert.Equal(t, GetHandType("23432", false), TwoPair)
	assert.Equal(t, GetHandType("A23A4", false), OnePair)
	assert.Equal(t, GetHandType("23456", false), HighCard)

	assert.Equal(t, GetHandType("JJJJ5", true), FiveOfAKind)
	assert.Equal(t, GetHandType("T55J5", true), FourOfAKind)
	assert.Equal(t, GetHandType("T55JJ", true), FourOfAKind)
	assert.Equal(t, GetHandType("JJ235", true), ThreeOfAKind)
	assert.Equal(t, GetHandType("1234J", true), OnePair)
}
