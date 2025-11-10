package day10

import "github.com/snopan/aoc-2023/src/helpers"

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
func (t Tile) ConnectedCoordinates(currCoord helpers.Coordinate) []helpers.Coordinate {

	x := currCoord.X
	y := currCoord.Y

	switch t {
	case Vertical:
		return []helpers.Coordinate{
			helpers.NewCoordinate(x, y-1),
			helpers.NewCoordinate(x, y+1),
		}
	case Horizontal:
		return []helpers.Coordinate{
			helpers.NewCoordinate(x-1, y),
			helpers.NewCoordinate(x+1, y),
		}
	case NorthEast:
		return []helpers.Coordinate{
			helpers.NewCoordinate(x, y-1),
			helpers.NewCoordinate(x+1, y),
		}
	case NorthWest:
		return []helpers.Coordinate{
			helpers.NewCoordinate(x, y-1),
			helpers.NewCoordinate(x-1, y),
		}
	case SouthWest:
		return []helpers.Coordinate{
			helpers.NewCoordinate(x, y+1),
			helpers.NewCoordinate(x-1, y),
		}
	case SouthEast:
		return []helpers.Coordinate{
			helpers.NewCoordinate(x, y+1),
			helpers.NewCoordinate(x+1, y),
		}
	case Start:
		return []helpers.Coordinate{
			helpers.NewCoordinate(x, y-1),
			helpers.NewCoordinate(x, y+1),
			helpers.NewCoordinate(x-1, y),
			helpers.NewCoordinate(x+1, y),
		}
	default:
		return []helpers.Coordinate{}
	}
}
