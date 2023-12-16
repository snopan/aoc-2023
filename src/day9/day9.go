package day9

import (
	"strings"

	"github.com/snopan/aoc-2023/src/helpers"
)

func SolutionPart1(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {

		ends := []int{}
		numbers := helpers.GetNumbersFromText(line)
		for !allZeroes(numbers) {
			ends = append(ends, numbers[len(numbers)-1])
			numbers = findDifferences(numbers)
		}

		extrapolate := 0
		for i := len(ends) - 1; i >= 0; i-- {
			extrapolate += ends[i]
		}

		total += extrapolate
	}

	return total
}

func SolutionPart2(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {

		starts := []int{}
		numbers := helpers.GetNumbersFromText(line)
		for !allZeroes(numbers) {
			starts = append(starts, numbers[0])
			numbers = findDifferences(numbers)
		}

		extrapolate := starts[len(starts)-1]
		for i := len(starts) - 2; i >= 0; i-- {
			extrapolate = starts[i] - extrapolate
		}

		total += extrapolate
	}

	return total
}

func findDifferences(numbers []int) []int {
	differences := []int{}
	prev := numbers[0]
	for _, n := range numbers[1:] {
		differences = append(differences, n-prev)
		prev = n
	}
	return differences
}

func allZeroes(numbers []int) bool {
	for _, n := range numbers {
		if n != 0 {
			return false
		}
	}
	return true
}
