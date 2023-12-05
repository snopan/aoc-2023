package helpers

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetNumbersFromText(text string) []int {
	re := regexp.MustCompile(`\d+`)
	numberTexts := re.FindAllStringSubmatch(text, -1)

	numbers := []int{}
	for _, n := range numberTexts {
		number, err := strconv.Atoi(n[0])
		if err != nil {
			panic(fmt.Sprintf("couldn't convert number %s", n))
		}
		numbers = append(numbers, number)
	}

	return numbers
}

func PrepareInput(input string) string {
	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "\n")
	input = strings.TrimSpace(input)
	return input
}
