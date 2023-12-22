package helpers

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func GetNumbersFromText(text string) []int {
	re := regexp.MustCompile(`-?\d+`)
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

func CombineIntArrToInt(numbers []int) int {
	res := 0
	op := 1
	for i := len(numbers) - 1; i >= 0; i-- {
		digits := numbers[i]
		res += numbers[i] * op
		op *= int(math.Pow(10, float64(DigitLen(digits))))
	}
	return res
}

func DigitLen(n int) int {
	count := 0
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}

func PrepareInput(input string) string {
	re := regexp.MustCompile(`\r?\n`)
	input = re.ReplaceAllString(input, "\n")
	input = strings.TrimSpace(input)
	return input
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Contains[K comparable](s []K, e K) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
