package day6

import (
	"math"
	"strings"

	"github.com/snopan/aoc-2023/src/helpers"
)

func SolutionPart1(input string) int {
	lines := strings.Split(input, "\n")
	times := helpers.GetNumbersFromText(lines[0])
	distances := helpers.GetNumbersFromText(lines[1])

	return calculateMarginError(times, distances)
}

func SolutionPart2(input string) int {
	lines := strings.Split(input, "\n")
	times := helpers.GetNumbersFromText(lines[0])
	distances := helpers.GetNumbersFromText(lines[1])

	times = []int{helpers.CombineIntArrToInt(times)}
	distances = []int{helpers.CombineIntArrToInt(distances)}

	return calculateMarginError(times, distances)
}

func calculateMarginError(times, distances []int) int {
	marginError := 1
	for i, time := range times {
		distToBeat := distances[i]

		var hold, travel int
		singleOption := false
		if math.Mod(float64(time), 2) == 0 {
			hold = time / 2
			travel = hold
			singleOption = true
		} else {
			hold = (time + 1) / 2
			travel = (time - 1) / 2
		}

		winningOptions := 0
		for hold != 0 {
			distance := hold * travel
			if distance > distToBeat {
				if singleOption {
					winningOptions++
					singleOption = false
				} else {
					winningOptions += 2
				}
				hold++
				travel--
			} else {
				break
			}
		}
		marginError *= winningOptions
	}

	return marginError
}
