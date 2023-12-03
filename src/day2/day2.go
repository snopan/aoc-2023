package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func SolutionPart1(input string) int {
	lines := strings.Split(input, "\n")

	actualBag := parseBag(lines[0])

	total := 0
	for _, game := range lines[1:] {
		if len(game) == 0 {
			continue
		}

		gameInfo := strings.Split(game, ":")
		gameID := getGameID(gameInfo[0])

		canFit := true
		rounds := strings.Split(gameInfo[1], ";")
		for _, round := range rounds {
			roundBag := parseBag(round)
			if canFit {
				canFit = actualBag.canFitBag(roundBag)
			}
		}

		if canFit {
			total += gameID
		}
	}

	return total
}

func SolutionPart2(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, game := range lines {
		if len(game) == 0 {
			continue
		}

		gameInfo := strings.Split(game, ":")

		var minimumBag *bag
		rounds := strings.Split(gameInfo[1], ";")
		for _, round := range rounds {
			roundBag := parseBag(round)
			if minimumBag == nil {
				minimumBag = roundBag
			} else {
				minimumBag.updateMaximum(roundBag)
			}
		}

		total += minimumBag.red * minimumBag.green * minimumBag.blue
	}

	return total
}

func getGameID(data string) int {
	split := strings.Split(data, " ")
	if len(split) != 2 {
		panic(fmt.Sprintf("this '%s' is not a valid input to get game id", data))
	}

	gameIDStr := split[1]
	gameID, err := strconv.Atoi(gameIDStr)
	if err != nil {
		panic("gameID found is not an integer")
	}

	return gameID
}

func parseBag(data string) *bag {
	bag := &bag{}

	if len(data) == 0 {
		return bag
	}

	for _, marbleLine := range strings.Split(data, ",") {
		marbleSplit := strings.Split(strings.TrimSpace(marbleLine), " ")
		marble := marbleSplit[1]

		amountStr := marbleSplit[0]
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			panic(fmt.Sprintf("amount for %s is not an integer", marble))
		}

		switch marble {
		case "red":
			bag.red = amount
		case "green":
			bag.green = amount
		case "blue":
			bag.blue = amount
		}
	}

	return bag
}
