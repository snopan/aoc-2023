package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	outputPath := os.Args[1]
	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic("cannot save ouptut")
	}

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		calibration := calibrate(line)
		total += calibration
	}

	outputFile.WriteString(fmt.Sprintf("%d", total))
}

func calibrate(data string) int {
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
