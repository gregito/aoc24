package day4

import (
	"fmt"
	"strings"

	"github.com/gregito/aoc24/utils"
)

const (
	xmas  = "XMAS"
	samx  = "SAMX"
	left  = true
	right = false
)

type Day4 struct{}

func (Day4) CompleteTask(input []string) {
	total := 0
	total += findAllHorizontal(input) + findAllVertical(input) + findAllDiagonal(input)

	fmt.Println("---------------")
	fmt.Printf("Number of XMAS: %d%s", total, utils.GetLineSeparator())
	fmt.Println("---------------")
}

func findAllHorizontal(input []string) int {
	count := 0
	for _, line := range input {
		count += strings.Count(line, xmas)
		count += strings.Count(line, samx)
	}
	return count
}

func findAllVertical(input []string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		line := ""
		for _, l := range input {
			line += l[i : i+1]
		}
		count += strings.Count(line, xmas)
		count += strings.Count(line, samx)
	}
	return count
}

func findAllDiagonal(input []string) int {
	return findAllHorizontal(rotate45Degrees(left, input)) + findAllHorizontal(rotate45Degrees(right, input))
}

func rotate45Degrees(toLeft bool, input []string) []string {
	n := len(input)
	splitInput := splitInput(input)
	lines := make([]string, 0)
	if toLeft {
		ctrl := n - 1
		for k := 0; k < 2*n-1; k++ {
			line := ""
			for i := 0; i < n; i++ {
				for j := n - 1; j >= 0; j-- {
					if j-i == ctrl {
						line += splitInput[i][j]
					}
				}
			}
			ctrl--
			lines = append(lines, line)
		}
	} else {
		ctrl := 0
		for k := 0; k < 2*n-1; k++ {
			line := ""
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					if i+j == ctrl {
						line += splitInput[i][j]
					}
				}
			}
			ctrl++
			lines = append(lines, line)
		}
	}
	return lines
}

func splitInput(input []string) [][]string {
	splitInput := make([][]string, 0)
	for _, s := range input {
		chars := make([]string, 0)
		for _, char := range strings.Split(s, "") {
			chars = append(chars, char)
		}
		splitInput = append(splitInput, chars)
	}
	return splitInput
}
