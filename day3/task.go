package day3

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gregito/aoc24/utils"
)

type Day3 struct{}

const (
	mainPattern            = "(mul\\(\\d*,\\d*\\))"
	conditionalMainPattern = "(mul\\(\\d*,\\d*\\))|(don\\'t\\(\\))|(do\\(\\))"
	numberPairPattern      = "\\d*,\\d*"
	numberSeparator        = ","
)

func (Day3) CompleteTask(input []string) {
	sum := getTotalSum(input)
	conditionalSum := getConditionalSum(input)
	fmt.Println("---------------")
	fmt.Printf("Sum: %d%s", sum, utils.GetLineSeparator())
	fmt.Printf("Conditional sum: %d%s", conditionalSum, utils.GetLineSeparator())
	fmt.Println("---------------")
}

func getTotalSum(input []string) int {
	sum := 0
	for _, line := range input {
		indices := regexp.MustCompile(mainPattern).FindAllStringIndex(line, -1)
		for i := range indices {
			core := line[indices[i][0]:indices[i][1]]
			numberIndices := regexp.MustCompile(numberPairPattern).FindAllStringIndex(core, -1)
			numExtract := strings.Split(core[numberIndices[0][0]:numberIndices[0][1]], numberSeparator)
			sum += utils.StringAsInt(numExtract[0]) * utils.StringAsInt(numExtract[1])
		}
	}
	return sum
}

func getConditionalSum(input []string) int {
	conditionalSum := 0
	line := ""
	for _, l := range input {
		line += l
	}
	indices := regexp.MustCompile(conditionalMainPattern).FindAllStringIndex(line, -1)
	skip := false
	for _, operation := range indices {
		action := line[operation[0]:operation[1]]
		if action == "do()" {
			skip = false
		} else if action == "don't()" {
			skip = true
		} else {
			numberIndices := regexp.MustCompile(numberPairPattern).FindAllStringIndex(action, -1)
			numExtract := strings.Split(action[numberIndices[0][0]:numberIndices[0][1]], numberSeparator)
			if !skip {
				conditionalSum += utils.StringAsInt(numExtract[0]) * utils.StringAsInt(numExtract[1])
			} else {
			}
		}
	}
	return conditionalSum
}
