package day04

import (
	"fmt"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type assignment struct {
	start, end int
}

func newAssignment(input []int) assignment {
	return assignment{
		start: input[0],
		end:   input[1],
	}
}

func (assignment assignment) contains(other assignment) bool {
	return assignment.start <= other.start && assignment.end >= other.end
}

func (assignment assignment) overlaps(other assignment) bool {
	return assignment.start <= other.end && assignment.end >= other.start
}

func parse(line string) (result [][]int) {
	result = make([][]int, 2)
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		panic(fmt.Errorf("bad input: %v", line))
	}
	for i := range parts {
		sections := strings.Split(parts[i], "-")
		if len(sections) != 2 {
			panic(fmt.Errorf("bad input: %v", line))
		}
		result[i] = []int{util.MustAtoi(sections[0]), util.MustAtoi(sections[1])}
	}
	return
}

func content(input <-chan string, full bool) (result int) {
	for line := range input {
		pair := parse(line)
		first := newAssignment(pair[0])
		second := newAssignment(pair[1])
		if full {
			if first.contains(second) || second.contains(first) {
				result++
			}
		} else {
			if first.overlaps(second) || second.overlaps(first) {
				result++
			}
		}
	}
	return
}

func A(input *challenge.Challenge) int {
	return content(input.Lines(), true)
}

func B(input *challenge.Challenge) int {
	return content(input.Lines(), false)
}
