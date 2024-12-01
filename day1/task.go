package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gregito/aoc24/utils"
)

const idSeparator = "   "

type Day1 struct{}

func (Day1) CompleteTask(input []string) {
	distance := 0
	similarityScore := 0
	idsOrdered := getIdsOrdered(input)
	rightColumn := getRightColumn(idsOrdered)
	for _, idPair := range idsOrdered {
		distance += calculateDistance(idPair)
		similarityScore += calculateSimilarity(idPair[0], rightColumn)
	}

	fmt.Println("---------------")
	fmt.Printf("Result ID distance: %d%s", distance, utils.GetLineSeparator())
	fmt.Printf("Result similarity score: %d%s", similarityScore, utils.GetLineSeparator())
	fmt.Println("---------------")
}

func calculateSimilarity(base int, rightColumn []int) int {
	similarity := 0
	for _, idFromRightrColumn := range rightColumn {
		if idFromRightrColumn == base {
			similarity++
		}
	}
	return base * similarity
}

func getRightColumn(ids [][]int) []int {
	rightColumn := make([]int, 0)
	for _, row := range ids {
		rightColumn = append(rightColumn, row[1])
	}
	return rightColumn
}

func calculateDistance(row []int) int {
	distance := row[0] - row[1]
	if distance < 0 {
		return distance * -1
	}
	return distance
}

func getIdsOrdered(input []string) [][]int {
	ids := make([][]int, 0)
	for _, idPair := range input {
		id := strings.Split(idPair, idSeparator)
		asInt := make([]int, 0)
		asInt = append(asInt, getAsInt(id[0]))
		asInt = append(asInt, getAsInt(id[1]))
		ids = append(ids, asInt)
	}
	orderIds(&ids)
	return ids
}

func orderIds(ids *[][]int) {
	var minLeftIndex int
	var minRightIndex int
	var leftTemp int
	var rightTemp int
	for i := 0; i < len(*ids)-1; i++ {
		minLeftIndex = i
		minRightIndex = i
		for j := i + 1; j < len(*ids); j++ {
			if (*ids)[j][0] < (*ids)[minLeftIndex][0] {
				minLeftIndex = j
			}
			if (*ids)[j][1] < (*ids)[minRightIndex][1] {
				minRightIndex = j
			}
		}
		leftTemp = (*ids)[i][0]
		rightTemp = (*ids)[i][1]
		(*ids)[i][0] = (*ids)[minLeftIndex][0]
		(*ids)[i][1] = (*ids)[minRightIndex][1]
		(*ids)[minLeftIndex][0] = leftTemp
		(*ids)[minRightIndex][1] = rightTemp
	}
}

func getAsInt(id string) int {
	asInt, _ := strconv.Atoi(id)
	// I mean c'mon, no error should occur
	return asInt
}
