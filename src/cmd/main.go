package main

import (
	"fmt"
	"io"
	"os"

	"github.com/snopan/aoc-2023/src/day1"
	"github.com/snopan/aoc-2023/src/day2"
	"github.com/snopan/aoc-2023/src/day3"
	"github.com/snopan/aoc-2023/src/day4"
	"github.com/snopan/aoc-2023/src/day5"
	"github.com/snopan/aoc-2023/src/day6"
	"github.com/snopan/aoc-2023/src/day7"
	"github.com/snopan/aoc-2023/src/day8"
	"github.com/snopan/aoc-2023/src/helpers"
)

func main() {
	day := os.Args[1]
	part := os.Args[2]
	inputBytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("cannot read input from stdin")
	}

	input := string(inputBytes)
	input = helpers.PrepareInput(input)
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
	case "4":
		switch part {
		case "1":
			solution = day4.SolutionPart1(input)
		case "2":
			solution = day4.SolutionPart2(input)
		}
	case "5":
		switch part {
		case "1":
			solution = day5.SolutionPart1(input)
		case "2":
			solution = day5.SolutionPart2(input)
		}
	case "6":
		switch part {
		case "1":
			solution = day6.SolutionPart1(input)
		case "2":
			solution = day6.SolutionPart2(input)
		}
	case "7":
		switch part {
		case "1":
			solution = day7.SolutionPart1(input)
		case "2":
			solution = day7.SolutionPart2(input)
		}
	case "8":
		switch part {
		case "1":
			solution = day8.SolutionPart1(input)
		case "2":
			solution = day8.SolutionPart2(input)
		}
	}

	fmt.Printf("Solution is: %+v", solution)
}
