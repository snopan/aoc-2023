package day8

import (
	"testing"

	"github.com/snopan/aoc-2023/src/helpers"
	"gotest.tools/assert"
)

var inputPart1 = helpers.PrepareInput(`
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`)

var inputPart2 = helpers.PrepareInput(`
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`)

func Test_SolutionPart1(t *testing.T) {
	solution := SolutionPart1(inputPart1)
	assert.Equal(t, solution, 2)
}

func Test_SolutionPart2(t *testing.T) {
	solution := SolutionPart2(inputPart2)
	assert.Equal(t, solution, 6)
}
