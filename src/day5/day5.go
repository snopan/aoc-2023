package day5

import (
	"fmt"
	"math"
	"strings"

	"github.com/snopan/aoc-2023/src/helpers"
)

func SolutionPart1(input string) int {
	data := strings.Split(input, "\n\n")
	initialSeeds := helpers.GetNumbersFromText(data[0])
	mappings := [][]*Conversion{}

	for _, mappingData := range data[1:] {
		mapping := parseMapping(mappingData)
		mappings = append(mappings, mapping)
	}

	lowestDest := math.MaxInt
	for _, seed := range initialSeeds {
		dest := seed
		for _, mapping := range mappings {
			dest = getDestValue(mapping, dest)
		}
		if dest < lowestDest {
			lowestDest = dest
		}
	}

	return lowestDest
}

func SolutionPart2(input string) int {
	data := strings.Split(input, "\n\n")
	initialSeeds := helpers.GetNumbersFromText(data[0])
	mappings := [][]*Conversion{}

	for _, mappingData := range data[1:] {
		mapping := parseMapping(mappingData)
		mappings = append(mappings, mapping)
	}

	initialSeedMapping := []*Bound{}
	for i := 0; i < len(initialSeeds); i += 2 {
		startSeed := initialSeeds[i]
		endSeed := startSeed + initialSeeds[i+1] - 1
		initialSeedMapping = append(initialSeedMapping, &Bound{
			start: startSeed,
			end:   endSeed,
		})
	}

	mappedBounds := initialSeedMapping
	for _, mapping := range mappings {
		unmappedBounds := mappedBounds
		mappedBounds = []*Bound{}

		for _, conversion := range mapping {
			newUnmappedBounds := []*Bound{}

			for _, bound := range unmappedBounds {
				mapped, unmapped := conversion.GetDestBounds(bound)
				if mapped != nil {
					mappedBounds = append(mappedBounds, mapped)
				}

				newUnmappedBounds = append(newUnmappedBounds, unmapped...)
			}
			unmappedBounds = newUnmappedBounds
		}
		mappedBounds = append(mappedBounds, unmappedBounds...)
	}

	lowestDest := math.MaxInt
	for _, bound := range mappedBounds {
		dest := bound.start
		if dest < lowestDest {
			lowestDest = dest
		}
	}

	return lowestDest
}

func getDestValue(mapping []*Conversion, source int) int {
	for _, conversion := range mapping {
		if conversion.ContainsSource(source) {
			return conversion.GetDest(source)
		}
	}
	return source
}

func parseMapping(data string) []*Conversion {
	m := []*Conversion{}

	lines := strings.Split(data, "\n")
	if len(lines) < 2 {
		return m
	}

	for _, line := range lines[1:] {
		values := helpers.GetNumbersFromText(line)
		if len(values) != 3 {
			panic(fmt.Sprintf("invalid map values %+v", values))
		}

		m = append(m, &Conversion{
			destStart:   values[0],
			sourceStart: values[1],
			rangeLen:    values[2],
		})
	}

	return m
}
