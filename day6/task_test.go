package day6

import (
	"log"
	"testing"

	"github.com/gregito/aoc24/utils"
)

func TestCompleteTask(t *testing.T) {
	lines, err := utils.ReadFileAndGetItsLines("d6_input.txt")
	if err != nil {
		log.Fatalf("%s%s\n", utils.TestSetupFailedMsg, err.Error())
	}
	Day6{}.CompleteTask(lines)
}
