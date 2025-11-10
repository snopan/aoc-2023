package day10

import (
	"math"
	"strings"

	"github.com/snopan/aoc-2023/src/helpers"
	"golang.org/x/exp/maps"
)

func SolutionPart1(input string) int {
	lines := strings.Split(input, "\n")

	tileMap := map[helpers.Coordinate]Tile{}
	var startCoordinate helpers.Coordinate
	for y, line := range lines {
		for x, t := range line {
			tile := ToTile(t)
			coordinate := helpers.NewCoordinate(x, y)
			tileMap[coordinate] = tile
			if tile == Start {
				startCoordinate = coordinate
			}
		}
	}

	travelledCoords := []helpers.Coordinate{}
	toProcessCoords := map[helpers.Coordinate][]helpers.Coordinate{
		startCoordinate: Start.ConnectedCoordinates(startCoordinate),
	}
	steps := 0
	for {
		newToProcessCoords := map[helpers.Coordinate][]helpers.Coordinate{}

		if len(maps.Keys(toProcessCoords)) == 0 {
			panic("no procesing left")
		}

		for coord, connectedCoords := range toProcessCoords {
			travelledCoords = append(travelledCoords, coord)
			for _, connectedCoord := range connectedCoords {
				nextTile, ok := tileMap[connectedCoord]
				if !ok || nextTile == Ground {
					continue
				}
				if helpers.Contains(travelledCoords, connectedCoord) {
					return steps
				}

				nextTileConnectedCoords := nextTile.ConnectedCoordinates(connectedCoord)

				isConnected := false
				newCoords := []helpers.Coordinate{}
				for _, ntcc := range nextTileConnectedCoords {
					if ntcc == coord {
						isConnected = true
					} else {
						newCoords = append(newCoords, ntcc)
					}
				}

				if isConnected {
					newToProcessCoords[connectedCoord] = newCoords
				}
			}
		}

		toProcessCoords = newToProcessCoords
		steps++
	}
}

func SolutionPart2(input string) int {
	lines := strings.Split(input, "\n")

	tileMap := map[helpers.Coordinate]Tile{}
	var startCoordinate helpers.Coordinate
	for y, line := range lines {
		for x, t := range line {
			tile := ToTile(t)
			coordinate := helpers.NewCoordinate(x, y)
			tileMap[coordinate] = tile
			if tile == Start {
				startCoordinate = coordinate
			}
		}
	}

	travelledCoords := []helpers.Coordinate{}
	toProcessCoords := map[helpers.Coordinate][]helpers.Coordinate{
		startCoordinate: Start.ConnectedCoordinates(startCoordinate),
	}
	done := false
	for !done {
		newToProcessCoords := map[helpers.Coordinate][]helpers.Coordinate{}

		if len(maps.Keys(toProcessCoords)) == 0 {
			panic("no procesing left")
		}

		for coord, connectedCoords := range toProcessCoords {
			travelledCoords = append(travelledCoords, coord)
			for _, connectedCoord := range connectedCoords {
				nextTile, ok := tileMap[connectedCoord]
				if !ok || nextTile == Ground {
					continue
				}
				if helpers.Contains(travelledCoords, connectedCoord) {
					done = true
					break
				}

				nextTileConnectedCoords := nextTile.ConnectedCoordinates(connectedCoord)

				isConnected := false
				newCoords := []helpers.Coordinate{}
				for _, ntcc := range nextTileConnectedCoords {
					if ntcc == coord {
						isConnected = true
					} else {
						newCoords = append(newCoords, ntcc)
					}
				}

				if isConnected {
					newToProcessCoords[connectedCoord] = newCoords
				}
			}
		}

		toProcessCoords = newToProcessCoords
	}

	insideCount := 0
	h := len(lines)
	w := len(lines[0])
	for y, line := range lines {
		for x := range line {
			if helpers.Contains(travelledCoords, helpers.NewCoordinate(x, y)) {
				continue
			}

			crosses := 0
			x2, y2 := x, y

			for x2 < w && y2 < h {
				coordinate := helpers.NewCoordinate(x2, y2)
				tile := tileMap[coordinate]
				if helpers.Contains(travelledCoords, coordinate) && tile != NorthEast && tile != SouthWest {
					crosses++
				}
				x2++
				y2++
			}

			if math.Mod(float64(crosses), 2) == 1 {
				insideCount++
			}
		}
	}
	return insideCount
}
