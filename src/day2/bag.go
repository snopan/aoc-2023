package day2

type bag struct {
	red   int
	green int
	blue  int
}

// This function returns whether the provided bag's contents
// can fit into this current bag
func (b *bag) canFitBag(otherBag *bag) bool {
	if b.red < otherBag.red {
		return false
	}
	if b.green < otherBag.green {
		return false
	}
	if b.blue < otherBag.blue {
		return false
	}
	return true
}

// Find the maxmimum for each marble in the current and provided bag
// then update current bag with that
func (b *bag) updateMaximum(otherBag *bag) {
	b.red = max(b.red, otherBag.red)
	b.green = max(b.green, otherBag.green)
	b.blue = max(b.blue, otherBag.blue)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
