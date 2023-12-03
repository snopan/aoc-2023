package main

import (
	"fmt"
	"io"
	"os"

	"github.com/snopan/aoc-2023/src/day1"
	"github.com/snopan/aoc-2023/src/day2"
	"github.com/snopan/aoc-2023/src/day3"
)

func main() {
	day := os.Args[1]
	part := os.Args[2]
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("cannot read input from stdin")
	}

	input := string(inputBytes)
	var solution any
	switch day {
	case "1":
		switch part {
		case "1":
			solution = day1.SolutionPart1(input)
		case "2":
			solution = day1.SolutionPart2(input)
		}
	case "2":
		switch part {
		case "1":
			solution = day2.SolutionPart1(input)
		case "2":
			solution = day2.SolutionPart2(input)
		}
	case "3":
		switch part {
		case "1":
			solution = day3.SolutionPart1(input)
		case "2":
			solution = day3.SolutionPart2(input)
		}
	}

	fmt.Printf("Solution is: %+v", solution)
}
