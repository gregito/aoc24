package day2

import (
	"fmt"
	"strings"

	"github.com/gregito/aoc24/utils"
)

type Day2 struct{}

const leveSeparator = " "

type report struct {
	levels []int
	safe   bool
}

func (Day2) CompleteTask(input []string) {
	safeReports := 0
	reports := readReports(input)
	for _, report := range reports {
		checkReport(&report)
		if report.safe {
			safeReports++
		}
	}
	fmt.Println("---------------")
	fmt.Printf("Number of safe reports: %d%s", safeReports, utils.GetLineSeparator())
	fmt.Println("---------------")
}

func readReports(input []string) []report {
	reports := make([]report, 0)
	for _, reportLine := range input {
		levels := make([]int, 0)
		for _, level := range strings.Split(reportLine, leveSeparator) {
			levels = append(levels, utils.StringAsInt(level))
		}
		reports = append(reports, report{levels: levels})
	}
	return reports
}

func checkReport(r *report) {
	safe := true
	shouldIncrease := true
	if r.levels[0] > r.levels[1] {
		shouldIncrease = false
	}
	for i := 0; i < len(r.levels)-1; i++ {
		if diff := r.levels[i+1] - r.levels[i]; shouldIncrease && !(diff > 0 && diff <= 3) {
			safe = false
			break
		}
		if diff := r.levels[i] - r.levels[i+1]; !shouldIncrease && !(diff > 0 && diff <= 3) {
			safe = false
			break
		}
	}
	r.safe = safe
}
