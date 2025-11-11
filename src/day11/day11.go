package day11

import (
	"strings"

	"github.com/snopan/aoc-2023/src/helpers"
)

func SolutionPart1(input string) int {
	lines := strings.Split(input, "\n")

	galaxies := []helpers.Coordinate{}

	// Keep track of rows and columns that have galaexies
	// Later when we calculate distance if a certain row and
	// column doesn't have galaxy we double the distance in
	// the relevant axis
	rowsHasGalaexies := map[int]bool{}
	columnHasGalaxies := map[int]bool{}

	for y, line := range lines {
		for x, t := range line {
			if t == '#' {
				galaxies = append(galaxies, helpers.NewCoordinate(x, y))
				rowsHasGalaexies[y] = true
				columnHasGalaxies[x] = true
			}
		}
	}

	distanceSum := 0
	galaxyPairs := generateCoordinatePairs(galaxies)
	for _, pair := range galaxyPairs {
		galaxyA := pair[0]
		galaxyB := pair[1]

		verticalDistance := 0
		if galaxyA.Y < galaxyB.Y {
			verticalDistance += calculateAxisDistance(galaxyA.Y, galaxyB.Y, rowsHasGalaexies)
		} else if galaxyA.Y > galaxyB.Y {
			verticalDistance += calculateAxisDistance(galaxyB.Y, galaxyA.Y, rowsHasGalaexies)
		}

		horizontalDistance := 0
		if galaxyA.X < galaxyB.X {
			horizontalDistance += calculateAxisDistance(galaxyA.X, galaxyB.X, columnHasGalaxies)
		} else if galaxyA.X > galaxyB.X {
			horizontalDistance += calculateAxisDistance(galaxyB.X, galaxyA.X, columnHasGalaxies)
		}
		distanceSum += verticalDistance + horizontalDistance
	}

	return distanceSum
}

func calculateAxisDistance(start int, end int, hasGalaxyOnAxis map[int]bool) int {
	distance := 0
	for i := start + 1; i <= end; i++ {
		if !hasGalaxyOnAxis[i] {
			distance += 2
		} else {
			distance++
		}
	}
	return distance
}

func generateCoordinatePairs(coordinates []helpers.Coordinate) [][]helpers.Coordinate {
	coordinatePairs := [][]helpers.Coordinate{}

	for index, coord := range coordinates {
		for i := index + 1; i < len(coordinates); i++ {
			pairedCoord := coordinates[i]
			coordinatePairs = append(coordinatePairs, []helpers.Coordinate{coord, pairedCoord})
		}
	}

	return coordinatePairs
}
