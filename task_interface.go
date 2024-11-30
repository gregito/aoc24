package aoc24

import "testing"

type Task interface {
	CompleteTask(input []string)
	TestCompleteTask(t *testing.T)
}
