package day3

import (
	"log"
	"testing"

	"github.com/gregito/aoc24/utils"
)

func TestCompleteTask(t *testing.T) {
	lines, err := utils.ReadFileAndGetItsLines("d3_input.txt")
	if err != nil {
		log.Fatalf("%s%s\n", utils.TestSetupFailedMsg, err.Error())
	}
	Day3{}.CompleteTask(lines)
}
