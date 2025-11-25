package day12

import (
	"strconv"
	"strings"
	"sync"
)

type procesStatus struct {
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

		workers := 10
		processes := []procesStatus{
			{
				processedSprings:                "",
				unprocessedSprings:              unprocessedSprings,
				currDamagedSpringsConnected:     0,
				expectedDamagedSpringsConnected: expectedDamagedSpringsConnected,
			},
		}
		var combinations []string
		var inflight int
		var poolLock sync.Mutex

		var wg sync.WaitGroup
		for w := 0; w < workers; w++ {
			wg.Add(1)
			go func(w int) {
				defer wg.Done()
				for {
					poolLock.Lock()
					if len(processes) == 0 {
						poolLock.Unlock()
						if inflight == 0 {
							return
						}
						continue
					}

					p := processes[0]
					processes = processes[1:]
					inflight++
					poolLock.Unlock()

					newProcesses, newCombinations := getSpringsCombinations(p)

					poolLock.Lock()
					processes = append(processes, newProcesses...)
					combinations = append(combinations, newCombinations...)
					inflight--
					poolLock.Unlock()
				}
			}(w)
		}

		wg.Wait()

		sum += len(combinations)
	}
	return sum
}

func getSpringsCombinations(status procesStatus) ([]procesStatus, []string) {
	if len(status.unprocessedSprings) == 0 {
		if status.currDamagedSpringsConnected > 0 {
			// If theres connected damaged springs then we should find that exact
			// amount of connected springs from other report (expectedConnectedSprings)
			if len(status.expectedDamagedSpringsConnected) != 1 ||
				status.currDamagedSpringsConnected != status.expectedDamagedSpringsConnected[0] {
				return []procesStatus{}, []string{}
			}
			return []procesStatus{}, []string{status.processedSprings}
		}

		// If theres no connected damaged springs and because it's at the end
		// there should be no expected connected damaged springs left
		if len(status.expectedDamagedSpringsConnected) > 0 {
			return []procesStatus{}, []string{}
		}
		return []procesStatus{}, []string{status.processedSprings}
	}

	currentSpring := status.unprocessedSprings[0]
	switch currentSpring {
	case '.':
		if status.currDamagedSpringsConnected > 0 {
			// If theres connected damaged springs then we should find that exact
			// amount of connected springs from other report (expectedConnectedSprings)
			if len(status.expectedDamagedSpringsConnected) == 0 ||
				status.currDamagedSpringsConnected != status.expectedDamagedSpringsConnected[0] {
				return []procesStatus{}, []string{}
			}

			// Mark the group of damaged springs connected as processed
			status.expectedDamagedSpringsConnected = status.expectedDamagedSpringsConnected[1:]
		}
		status.processedSprings += "."
		status.unprocessedSprings = status.unprocessedSprings[1:]

		// Reset the counter for damaged springs connected
		status.currDamagedSpringsConnected = 0
		return []procesStatus{status}, []string{}
	case '#':
		status.processedSprings += "#"
		status.unprocessedSprings = status.unprocessedSprings[1:]
		status.currDamagedSpringsConnected++
		return []procesStatus{status}, []string{}
	case '?':
		damagedSpringStatus := procesStatus{
			processedSprings:                status.processedSprings,
			unprocessedSprings:              "#" + status.unprocessedSprings[1:],
			currDamagedSpringsConnected:     status.currDamagedSpringsConnected,
			expectedDamagedSpringsConnected: status.expectedDamagedSpringsConnected,
		}
		workingSpringStatus := procesStatus{
			processedSprings:                status.processedSprings,
			unprocessedSprings:              "." + status.unprocessedSprings[1:],
			currDamagedSpringsConnected:     status.currDamagedSpringsConnected,
			expectedDamagedSpringsConnected: status.expectedDamagedSpringsConnected,
		}
		return []procesStatus{damagedSpringStatus, workingSpringStatus}, []string{}
	default:
		panic("unknown spring value")
	}
}
