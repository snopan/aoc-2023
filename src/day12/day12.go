package day12

import (
	"strconv"
	"strings"
)

type ProcesStatus struct {
	processedSprings                string
	unprocessedSprings              string
	currDamagedSpringsConnected     int
	expectedDamagedSpringsConnected []int
}

func SolutionPart1(input string) int {
	lines := strings.Split(input, "\n")

	var sum int
	for _, line := range lines {
		parts := strings.Split(line, " ")
		unprocessedSprings := parts[0]
		var expectedDamagedSpringsConnected []int
		for _, connectedSpringsStr := range strings.Split(parts[1], ",") {
			connectedSprings, err := strconv.Atoi(connectedSpringsStr)
			if err != nil {
				panic(err)
			}
			expectedDamagedSpringsConnected = append(expectedDamagedSpringsConnected, connectedSprings)
		}

		springsCombinations := getSpringsCombinations(ProcesStatus{
			processedSprings:                "",
			unprocessedSprings:              unprocessedSprings,
			currDamagedSpringsConnected:     0,
			expectedDamagedSpringsConnected: expectedDamagedSpringsConnected,
		})

		sum += len(springsCombinations)
	}
	return sum
}

func getSpringsCombinations(status ProcesStatus) []string {
	if len(status.unprocessedSprings) == 0 {
		if status.currDamagedSpringsConnected > 0 {
			// If theres connected damaged springs then we should find that exact
			// amount of connected springs from other report (expectedConnectedSprings)
			if len(status.expectedDamagedSpringsConnected) != 1 ||
				status.currDamagedSpringsConnected != status.expectedDamagedSpringsConnected[0] {
				return []string{}
			}
			return []string{status.processedSprings}
		}

		// If theres no connected damaged springs and because it's at the end
		// there should be no expected connected damaged springs left
		if len(status.expectedDamagedSpringsConnected) > 0 {
			return []string{}
		}
		return []string{status.processedSprings}
	}

	currentSpring := status.unprocessedSprings[0]
	switch currentSpring {
	case '.':
		if status.currDamagedSpringsConnected > 0 {
			// If theres connected damaged springs then we should find that exact
			// amount of connected springs from other report (expectedConnectedSprings)
			if len(status.expectedDamagedSpringsConnected) == 0 ||
				status.currDamagedSpringsConnected != status.expectedDamagedSpringsConnected[0] {
				return []string{}
			}

			// Mark the group of damaged springs connected as processed
			status.expectedDamagedSpringsConnected = status.expectedDamagedSpringsConnected[1:]
		}
		status.processedSprings += "."
		status.unprocessedSprings = status.unprocessedSprings[1:]

		// Reset the counter for damaged springs connected
		status.currDamagedSpringsConnected = 0
		return getSpringsCombinations(status)
	case '#':
		status.processedSprings += "#"
		status.unprocessedSprings = status.unprocessedSprings[1:]
		status.currDamagedSpringsConnected++
		return getSpringsCombinations(status)
	case '?':
		var springsCombinations []string
		status.unprocessedSprings = "#" + status.unprocessedSprings[1:]
		springsCombinations = append(springsCombinations, getSpringsCombinations(status)...)
		status.unprocessedSprings = "." + status.unprocessedSprings[1:]
		springsCombinations = append(springsCombinations, getSpringsCombinations(status)...)
		return springsCombinations
	default:
		panic("unknown spring value")
	}
}
