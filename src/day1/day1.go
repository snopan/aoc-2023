package day1

import (
	"fmt"
	"strings"
	"unicode"
)

func SolutionPart1(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		calibration := calibrate(line)
		total += calibration
	}

	return total
}

func SolutionPart2(input string) int {
	lines := strings.Split(input, "\n")

	total := 0
	for _, line := range lines {
		convertedLine := convertWordNumbers(line)
		calibration := calibrate(convertedLine)
		total += calibration
	}

	return total
}

func calibrate(data string) int {
	if len(data) == 0 {
		return 0
	}

	front := -1
	back := -1
	for _, s := range data {
		if unicode.IsDigit(s) {
			if front == -1 {
				front = int(s - '0')
			}
			back = int(s - '0')
		}
	}

	return front*10 + back
}

func convertWordNumbers(data string) string {
	wordNumberTree := generateWordNumberTree()
	curr := wordNumberTree

	parsed := ""
	checking := ""
	var ok bool
	for _, s := range data {
		letter := string(s)

		// Check if this letter continues on this number word
		curr, ok = curr.children[letter]
		if !ok {

			// When it doesn't the add all the previouis checking to parsed and reset
			curr = wordNumberTree
			parsed += checking
			checking = ""

			// While this letter is not valid for the current checking number
			// it may be a start of a different number so check again
			// if it's not then it's safe to add this letter to parsed
			curr, ok = curr.children[letter]
			if !ok {
				curr = wordNumberTree
				parsed += letter
			} else {
				checking += letter
			}
		} else {
			checking += letter
		}

		if curr.validValue() {
			parsed += fmt.Sprintf("%d", curr.value)
			curr = wordNumberTree
			checking = ""
		}
	}
	parsed += checking

	return parsed
}

func generateWordNumberTree() *node {
	root := &node{
		value:    -1,
		children: map[string]*node{},
	}
	addWordNumberNodes(root, "one", 1)
	addWordNumberNodes(root, "two", 2)
	addWordNumberNodes(root, "three", 3)
	addWordNumberNodes(root, "four", 4)
	addWordNumberNodes(root, "five", 5)
	addWordNumberNodes(root, "six", 6)
	addWordNumberNodes(root, "seven", 7)
	addWordNumberNodes(root, "eight", 8)
	addWordNumberNodes(root, "nine", 9)
	return root
}

func addWordNumberNodes(root *node, word string, number int) {
	curr := root
	value := -1
	for i, letter := range word {
		if i == len(word)-1 {
			value = number
		}
		curr = curr.addNode(string(letter), value)
	}
}
