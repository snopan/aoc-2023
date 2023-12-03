package day3

import "fmt"

type Grid struct {
	grid map[int]map[int]*int
}

func newGrid() *Grid {
	return &Grid{
		grid: make(map[int]map[int]*int),
	}
}

func (g *Grid) addPoint(x, y, value int) {
	column, ok := g.grid[x]
	if !ok {
		column = make(map[int]*int)
		g.grid[x] = column
	}

	_, ok = column[y]
	if ok {
		panic("point already exists")
	}
	column[y] = &value
}

// Remove number from grid, given the x and y
// where this number is at, try to move left and right
// in columns to remove this number
func (g *Grid) removeNumber(x, y, value int) {
	g.grid[x][y] = nil
	g.loopRemoveNumber(x, y, value, 1)
	g.loopRemoveNumber(x, y, value, -1)
}

func (g *Grid) loopRemoveNumber(x, y, value, moveBy int) {
	i := moveBy
	for {
		v := g.grid[x][y+i]
		if v == nil {
			break
		} else if *v != value {
			panic(fmt.Sprintf("while moving by %d to remove number %d, found a different number %d", moveBy, value, *v))
		}

		g.grid[x][y+i] = nil
		i += moveBy
	}
}

func (g *Grid) getPoint(x, y int) (int, bool) {
	value := g.grid[x][y]
	if value == nil {
		return 0, false
	}
	return *value, true
}

type Point struct {
	x int
	y int
}

func (p *Point) getAdjacentPoints() []*Point {
	adjacentPoints := []*Point{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i != 0 || j != 0 {
				adjacentPoints = append(adjacentPoints, &Point{
					x: p.x + i,
					y: p.y + j,
				})
			}
		}
	}
	return adjacentPoints
}
