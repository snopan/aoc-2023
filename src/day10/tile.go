package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type Tile int

const (
	Vertical Tile = iota
	Horizontal
	NorthEast
	NorthWest
	SouthWest
	SouthEast
	Ground
	Start
)

func ToTile(t rune) Tile {
	switch t {
	case '|':
		return Vertical
	case '-':
		return Horizontal
	case 'L':
		return NorthEast
	case 'J':
		return NorthWest
	case '7':
		return SouthWest
	case 'F':
		return SouthEast
	case '.':
		return Ground
	case 'S':
		return Start
	default:
		panic("invalid tile")
	}
}

// Given the current coordinate and the type of tile this is
// get the coordinates that this tile will be connected to
func (t Tile) ConnectedCoordinates(currCoord string) []string {

	x, y := FromCoordinate(currCoord)

	switch t {
	case Vertical:
		return []string{
			Coordinate(x, y-1),
			Coordinate(x, y+1),
		}
	case Horizontal:
		return []string{
			Coordinate(x-1, y),
			Coordinate(x+1, y),
		}
	case NorthEast:
		return []string{
			Coordinate(x, y-1),
			Coordinate(x+1, y),
		}
	case NorthWest:
		return []string{
			Coordinate(x, y-1),
			Coordinate(x-1, y),
		}
	case SouthWest:
		return []string{
			Coordinate(x, y+1),
			Coordinate(x-1, y),
		}
	case SouthEast:
		return []string{
			Coordinate(x, y+1),
			Coordinate(x+1, y),
		}
	case Start:
		return []string{
			Coordinate(x, y-1),
			Coordinate(x, y+1),
			Coordinate(x-1, y),
			Coordinate(x+1, y),
		}
	default:
		return []string{}
	}
}

func Coordinate(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func FromCoordinate(c string) (int, int) {
	coords := strings.Split(c, "-")

	x, err := strconv.Atoi(coords[0])
	if err != nil {
		panic("can't parse x from coordinate string")
	}
	y, err := strconv.Atoi(coords[1])
	if err != nil {
		panic("can't parse y from coordinate string")
	}

	return x, y
}
