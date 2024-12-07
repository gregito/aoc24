package day6

import (
	"fmt"
	"strings"

	"github.com/gregito/aoc24/utils"
)

type Day6 struct{}

const (
	leftDirection  = "<"
	rightDirection = ">"
	upDirection    = "^"
	downDirection  = "v"
	obstacle       = "#"
)

func (Day6) CompleteTask(input []string) {
	visitedFields := getVisitedFields(createMap(input))

	fmt.Println("---------------")
	fmt.Printf("Number of visited fields: %d%s", len(visitedFields), utils.GetLineSeparator())
	fmt.Println("---------------")
}

func getVisitedFields(guardMap [][]string) map[string]bool {
	outOfBound := false
	guardI, guardJ := getStartIndex(guardMap)
	visitedFields := make(map[string]bool)
	visitedFields[createMapEntryFromPosition(guardI, guardJ)] = true
	for !outOfBound {
		if guardMap[guardI][guardJ] == upDirection {
			for guardI != -1 {
				if guardI == 0 {
					outOfBound = true
					break
				}
				if guardMap[guardI-1][guardJ] == obstacle {
					guardMap[guardI][guardJ] = rightDirection
					break
				} else {
					guardI--
					visitedFields[createMapEntryFromPosition(guardI, guardJ)] = true
				}
			}
		}
		if guardMap[guardI][guardJ] == downDirection && !outOfBound {
			for guardI != len(guardMap) {
				if guardI == len(guardMap)-1 {
					outOfBound = true
					break
				}
				if guardMap[guardI+1][guardJ] == obstacle {
					guardMap[guardI][guardJ] = leftDirection
					break
				} else {
					guardI++
					visitedFields[createMapEntryFromPosition(guardI, guardJ)] = true
				}
			}
		}
		if guardMap[guardI][guardJ] == rightDirection && !outOfBound {
			for guardJ != len(guardMap[guardI])-1 {
				if guardJ == len(guardMap[guardI])-1 {
					outOfBound = true
					break
				}
				if guardMap[guardI][guardJ+1] == obstacle {
					guardMap[guardI][guardJ] = downDirection
					break
				} else {
					guardJ++
					visitedFields[createMapEntryFromPosition(guardI, guardJ)] = true
				}
			}
		}
		if guardMap[guardI][guardJ] == leftDirection && !outOfBound {
			for guardJ != -1 {
				if guardJ == 0 {
					outOfBound = true
					break
				}
				if guardMap[guardI][guardJ-1] == obstacle {
					guardMap[guardI][guardJ] = upDirection
					break
				} else {
					guardJ--
					visitedFields[createMapEntryFromPosition(guardI, guardJ)] = true
				}
			}
		}
	}
	return visitedFields
}

func getStartIndex(guardMap [][]string) (int, int) {
	for i, row := range guardMap {
		for j, field := range row {
			if field == upDirection {
				return i, j
			}
		}
	}
	return -1, -1
}

func createMap(input []string) [][]string {
	guardMap := make([][]string, 0)
	for _, line := range input {
		split := strings.Split(line, "")
		row := make([]string, 0)
		for _, field := range split {
			row = append(row, field)
		}
		guardMap = append(guardMap, row)
	}
	return guardMap
}

func createMapEntryFromPosition(i int, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}
