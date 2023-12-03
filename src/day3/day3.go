package day3

import (
	"math"
	"strings"
	"unicode"
)

func SolutionPart1(input string) int {
	numberGrid, symbolPoints := buildNumberGridAndSymbolPoints(input)

	// loop through symbols and find part numbers and add it to total
	total := 0
	for _, point := range symbolPoints {
		adjacentPoints := point.getAdjacentPoints()
		for _, ap := range adjacentPoints {
			number, ok := numberGrid.getPoint(ap.x, ap.y)
			if ok {
				total += number

				// Once number is used we can remove it
				numberGrid.removeNumber(ap.x, ap.y, number)
			}
		}
	}

	return total
}

func SolutionPart2(input string) int {
	numberGrid, symbolPoints := buildNumberGridAndSymbolPoints(input)

	// loop through symbols and only when we find exactly two part numbers, then add it to total
	total := 0
	for _, point := range symbolPoints {
		gearRatio := 1
		numParts := 0
		adjacentPoints := point.getAdjacentPoints()
		for _, ap := range adjacentPoints {
			number, ok := numberGrid.getPoint(ap.x, ap.y)
			if ok {
				numParts++
				gearRatio *= number

				// Once number is used we can remove it
				numberGrid.removeNumber(ap.x, ap.y, number)
			}
		}

		if numParts == 2 {
			total += gearRatio
		} else {
			numberGrid.undoRecentRemovals()
		}
	}

	return total
}

func buildNumberGridAndSymbolPoints(data string) (*Grid, []*Point) {
	numberGrid := newGrid()
	symbolPoints := []*Point{}

	// Build number grid and symbol points
	rows := strings.Split(data, "\n")
	for x, row := range rows {

		digits := []int{}
		yToAdd := []int{}
		for y, char := range strings.TrimSpace(row) {

			// When it's part of a number, keep track of it to add to grid later
			if unicode.IsDigit(char) {
				digits = append(digits, int(char-'0'))
				yToAdd = append(yToAdd, y)
			} else {

				// Add the number to our grid because it's finished
				if len(yToAdd) != 0 {
					number := digitsToNumber(digits)
					for _, ys := range yToAdd {
						numberGrid.addPoint(x, ys, number)
					}
					digits = []int{}
					yToAdd = []int{}
				}

				// For any character that is not '.' and digit it will be a symbol
				if char != '.' {
					symbolPoints = append(symbolPoints, &Point{x, y})
				}
			}
		}

		// If we're still have a number at the end, add that number to our grid
		if len(yToAdd) != 0 {
			number := digitsToNumber(digits)
			for _, ys := range yToAdd {
				numberGrid.addPoint(x, ys, number)
			}
		}
	}

	return numberGrid, symbolPoints
}

func digitsToNumber(digits []int) int {
	number := 0
	power := 0
	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		number += int(math.Pow(10, float64(power))) * digit
		power++
	}
	return number
}
