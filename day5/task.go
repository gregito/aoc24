package day5

import (
	"fmt"
	"strings"

	"github.com/gregito/aoc24/utils"
)

type Day5 struct{}

type update struct {
	pages        []int
	inRightOrder bool
}

type rule struct {
	firstNumber  int
	secondNumber int
}

func (Day5) CompleteTask(input []string) {
	updates := getUpdates(input)

	validUpdates := checkRules(&updates, input)
	validUpdateSum := 0
	for _, validUpdate := range validUpdates {
		validUpdateSum += validUpdate[(len(validUpdate) / 2)]
	}

	fmt.Println("---------------")
	fmt.Printf("Sum of valid page mids: %d%s", validUpdateSum, utils.GetLineSeparator())
	fmt.Println("---------------")
}

func checkRules(updates *[]update, input []string) [][]int {
	validUpdates := make([][]int, 0)
	rules := getRulesFromInput(input)
	for _, update := range *updates {
		valid := true
		for _, rule := range rules {
			if sliceContainsBothNumbers(update.pages, rule.firstNumber, rule.secondNumber) {
				indexOfFirst := 0
				indexOfSecond := 0
				for i, page := range update.pages {
					if page == rule.firstNumber {
						indexOfFirst = i
						continue
					}
					if page == rule.secondNumber {
						indexOfSecond = i
					}
				}
				if indexOfFirst >= indexOfSecond {
					valid = false
				}
			}
		}
		if valid {
			validUpdates = append(validUpdates, update.pages)
		}
	}
	return validUpdates
}

func getRulesFromInput(input []string) []rule {
	rules := make([]rule, 0)
	for _, line := range input {
		if line == "" {
			break
		}
		split := strings.Split(line, "|")
		firstNumber := utils.StringAsInt(split[0])
		secondNumber := utils.StringAsInt(split[1])
		rules = append(rules, rule{
			firstNumber:  firstNumber,
			secondNumber: secondNumber,
		})
	}
	return rules
}

func getUpdates(input []string) []update {
	updates := make([]update, 0)
	for i := updateStartIndex(input); i < len(input); i++ {
		line := strings.Split(input[i], ",")
		updatePages := make([]int, 0)
		for _, num := range line {
			updatePages = append(updatePages, utils.StringAsInt(num))
		}
		updates = append(updates, update{pages: updatePages})
	}
	return updates
}

func sliceContainsBothNumbers(updatePages []int, firstNumber int, secondNumber int) bool {
	count := 0
	for i := 0; i < len(updatePages) && count < 2; i++ {
		if updatePages[i] == firstNumber || updatePages[i] == secondNumber {
			count++
		}
	}
	return count == 2
}

func updateStartIndex(input []string) int {
	for i, _ := range input {
		if len(input[i]) == 0 {
			return i + 1
		}
	}
	return -1 // Thus will break as out of bound
}
