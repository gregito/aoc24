package day5

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gregito/aoc24/utils"
)

var rules = make([]rule, 0)

type Day5 struct{}

type update struct {
	pages []int
}

type rule struct {
	firstNumber  int
	secondNumber int
}

func (Day5) CompleteTask(input []string) {
	rules = getRulesFromInput(input)
	updates := getUpdates(input)

	validUpdates := checkRules(&updates)
	validUpdateSum := getSumOfMidValues(validUpdates)

	collectUnorderedUpdates(updates, validUpdates)
	orderedSum := 0

	fmt.Println("---------------")
	fmt.Printf("Sum of valid page mids: %d%s", validUpdateSum, utils.GetLineSeparator())
	fmt.Printf("Sum of newly ordered page mids: %d%s", orderedSum, utils.GetLineSeparator())
	fmt.Println("---------------")
}

func reorder(upd update) update {
	//ordered := false
	for i := 0; i < len(upd.pages)-1; i++ {
		ordered := false
		for j, rule := range rules {
			
		}
	}
	return upd
}

func getSumOfMidValues(updates []update) int {
	updateSum := 0
	for _, u := range updates {
		updateSum += u.pages[(len(u.pages) / 2)]
	}
	return updateSum
}

func checkRules(updates *[]update) []update {
	validUpdates := make([]update, 0)
	for _, update := range *updates {
		valid := true
		for _, rule := range rules {
			valid = checkUpdateWithRule(update, rule, valid)
		}
		if valid {
			validUpdates = append(validUpdates, update)
		}
	}
	return validUpdates
}

func checkUpdateWithRule(update update, rule rule, valid bool) bool {
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
	return valid
}

func collectUnorderedUpdates(allUpdates []update, orderedUpdates []update) []update {
	unorderedOnes := make([]update, 0)
	for _, update := range allUpdates {
		found := false
		for _, orderedUpdate := range orderedUpdates {
			if reflect.DeepEqual(update.pages, orderedUpdate.pages) {
				found = true
				break
			}
		}
		if !found {
			unorderedOnes = append(unorderedOnes, update)
		}
	}
	return unorderedOnes
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
