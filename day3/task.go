package day3

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gregito/aoc24/utils"
)

type Day3 struct{}

const (
	mainPattern       = "(mul\\(\\d*,\\d*\\))"
	numberPairPattern = "\\d*,\\d*"
	numberSeparator   = ","
)

func (Day3) CompleteTask(input []string) {
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
	fmt.Println("---------------")
	fmt.Printf("Sum: %d%s", sum, utils.GetLineSeparator())
	fmt.Println("---------------")
}
