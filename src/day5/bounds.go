package day5

type Bound struct {
	start int
	end   int
}

type Conversion struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

func (c *Conversion) ContainsSource(source int) bool {
	return source >= c.sourceStart && source <= c.SourceEnd()
}

func (c *Conversion) SourceEnd() int {
	return c.sourceStart + c.rangeLen - 1
}

func (c *Conversion) GetDest(source int) int {
	diff := source - c.sourceStart
	return c.destStart + diff
}

func (c *Conversion) GetDestBounds(b *Bound) (*Bound, []*Bound) {
	containsStart := c.ContainsSource(b.start)
	containsEnd := c.ContainsSource(b.end)

	outsideSourceStartBound := &Bound{
		start: b.start,
		end:   c.sourceStart - 1,
	}
	outsideSourceEndBound := &Bound{
		start: c.SourceEnd() + 1,
		end:   b.end,
	}

	// The bound provided is within source start and end
	if containsStart && containsEnd {
		return &Bound{
			start: c.GetDest(b.start),
			end:   c.GetDest(b.end),
		}, []*Bound{}

		// The bound has it's end inside the source range but it's start is outside
	} else if !containsStart && containsEnd {
		return &Bound{
				start: c.GetDest(c.sourceStart),
				end:   c.GetDest(b.end),
			},
			[]*Bound{outsideSourceStartBound}

		// The bound has it's start inside the source range but it's end is outside
	} else if containsStart && !containsEnd {
		return &Bound{
				start: c.GetDest(b.start),
				end:   c.GetDest(c.SourceEnd()),
			},
			[]*Bound{outsideSourceEndBound}

		// The bound is larger than the source range
	} else if b.start < c.sourceStart && b.end > c.SourceEnd() {
		return &Bound{
				start: c.GetDest(c.sourceStart),
				end:   c.GetDest(c.SourceEnd()),
			},
			[]*Bound{
				outsideSourceStartBound,
				outsideSourceEndBound,
			}

		// The bound is fully not part of the source range
	} else {
		return nil, []*Bound{b}
	}
}
