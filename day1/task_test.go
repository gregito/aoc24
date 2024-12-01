package day1

import (
	"log"
	"testing"

	"github.com/gregito/aoc24/utils"
)

func TestCompleteTask(t *testing.T) {
	lines, err := utils.ReadFileAndGetItsLines("d1_input.txt")
	if err != nil {
		log.Fatalf("%s%s\n", utils.TestSetupFailedMsg, err.Error())
	}
	Day1{}.CompleteTask(lines)

	// 1970720
}
