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

// A bit cursed
func convertWordNumbers(data string) string {
	wordNumberTree := generateWordNumberTree()

	var checkingTree *node
	var ok bool
	var checkingStartIndex int
	tempStartLetter := false
	parsed := ""
	checking := ""
	for i := 0; i < len(data); i++ {
		letter := string(data[i])

		// If we're not checking any word at the moment, see if this
		// letter is the start of a word
		if checking == "" {
			checkingTree, ok = wordNumberTree.children[letter]

			// If it is then mark where we started and add this letter to current checking
			if ok {
				checkingStartIndex = i
				checking += letter

				// If it's not a valid start then add this letter to parsed
			} else {
				parsed += letter
			}

			// If we are checking then see if this continues the word
		} else {
			checkingTree, ok = checkingTree.children[letter]

			// If it continues add this letter to current checking
			if ok {
				checking += letter

				// The letter is at the end of the word so add the value and reset
				if checkingTree.validValue() {
					parsed += fmt.Sprintf("%d", checkingTree.value)
					checking = ""

					// This last letter for the word may also be teh start of another word
					checkingTree, ok = wordNumberTree.children[letter]
					if ok {
						checkingStartIndex = i
						checking += letter
						tempStartLetter = true
					}
				}

				// If it's not a valid word at the end then add the start letter of
				// this word to parsed and move index back to one after the start
			} else {

				// If the start letter used is the end of a previous word,
				// we don't need to at that start letter back in
				if !tempStartLetter {
					startLetter := string(data[checkingStartIndex])
					parsed += startLetter
				}

				tempStartLetter = false
				checking = ""
				i = checkingStartIndex
			}
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
