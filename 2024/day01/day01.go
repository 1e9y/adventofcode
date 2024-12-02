package day01

import (
	"slices"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func parse(input []string) (left, right []int) {
	for _, line := range input {
		parts := strings.Split(line, "   ")
		left = append(left, util.MustAtoi(parts[0]))
		right = append(right, util.MustAtoi(parts[1]))
	}

	return
}

func distance(input []string) (answer int) {
	left, right := parse(input)

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		answer += util.AbsInt(left[i] - right[i])
	}

	return
}

func similarity(input []string) (answer int) {
	left, right := parse(input)
	freq := map[int]int{}

	for _, r := range right {
		freq[r]++
	}

	for _, l := range left {
		answer += l * freq[l]
	}

	return
}

func A(input *challenge.Challenge) int {
	return distance(input.LineSlice())
}

func B(input *challenge.Challenge) int {
	return similarity(input.LineSlice())
}
