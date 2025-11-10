package helpers

type Coordinate struct {
	X int
	Y int
}

func NewCoordinate(x int, y int) Coordinate {
	return Coordinate{
		X: x,
		Y: y,
	}
}
