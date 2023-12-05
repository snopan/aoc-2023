package day4

import (
	"math"
	"strings"

	"github.com/snopan/aoc-2023/src/helpers"
)

func SolutionPart1(input string) int {
	games := strings.Split(input, "\n")

	total := 0
	for _, game := range games {
		game = strings.TrimSpace(game)
		if len(game) == 0 {
			continue
		}

		cardInfo := strings.Split(game, ":")[1]
		cards := strings.Split(cardInfo, "|")
		winningNumbersStr := cards[0]
		ownNumbersStr := cards[1]

		winningNumbers := helpers.GetNumbersFromText(winningNumbersStr)
		winningNumbersMap := numberArrToMap(winningNumbers)
		ownNumbers := helpers.GetNumbersFromText(ownNumbersStr)

		numMatches := numberMatches(winningNumbersMap, ownNumbers)
		total += calculateScore(numMatches)
	}

	return total
}

func SolutionPart2(input string) int {
	games := strings.Split(input, "\n")

	total := 0
	cardCopies := []int{}
	for _, game := range games {
		game = strings.TrimSpace(game)
		if len(game) == 0 {
			continue
		}

		cardInfo := strings.Split(game, ":")[1]
		cards := strings.Split(cardInfo, "|")
		winningNumbersStr := cards[0]
		ownNumbersStr := cards[1]

		winningNumbers := helpers.GetNumbersFromText(winningNumbersStr)
		winningNumbersMap := numberArrToMap(winningNumbers)
		ownNumbers := helpers.GetNumbersFromText(ownNumbersStr)

		numMatches := numberMatches(winningNumbersMap, ownNumbers)
		currCopies := 1
		if len(cardCopies) > 0 {
			currCopies += cardCopies[0]
			cardCopies = cardCopies[1:]
		}
		total += currCopies

		if numMatches > 0 {
			currLen := len(cardCopies)
			for i := 0; i < min(numMatches, currLen); i++ {
				cardCopies[i] += currCopies
			}
			if numMatches > currLen {
				for i := 0; i < numMatches-currLen; i++ {
					cardCopies = append(cardCopies, currCopies)
				}
			}
		}
	}

	return total
}

func numberMatches(winningNumbersMap map[int]bool, ownNumbers []int) int {
	num := 0
	for _, n := range ownNumbers {
		_, ok := winningNumbersMap[n]
		if ok {
			num++
		}
	}
	return num
}

func calculateScore(numMatches int) int {
	if numMatches > 1 {
		return int(math.Pow(2, (float64(numMatches) - 1)))
	} else {
		return numMatches
	}
}

func numberArrToMap(numberArr []int) map[int]bool {
	numberMap := map[int]bool{}
	for _, n := range numberArr {
		numberMap[n] = true
	}
	return numberMap
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
