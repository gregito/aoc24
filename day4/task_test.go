package day4

import (
	"fmt"
	"log"
	"testing"

	"github.com/gregito/aoc24/utils"
)

func TestCompleteTask(t *testing.T) {
	lines, err := utils.ReadFileAndGetItsLines("d4_input.txt")
	if err != nil {
		log.Fatalf("%s%s\n", utils.TestSetupFailedMsg, err.Error())
	}
	Day4{}.CompleteTask(lines)
}

func TestRotate(t *testing.T) {
	input := []string{
		"ABCD",
		"EFGH",
		"IJKL",
		"MNOP",
	}
	result := rotate45Degrees(true, input)

	fmt.Println(result)

	/*
		D
		CH
		BGL
		AFKP
		EJO
		IN
		M
	*/
}
