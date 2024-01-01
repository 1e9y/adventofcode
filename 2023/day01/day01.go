package day01

import (
	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
	"sort"
)

func calories(input <-chan string) []int {
	elfs := make([]int, 0)
	sum := 0
	for c := range input {
		if c == "" {
			elfs = append(elfs, sum)
			sum = 0
		} else {
			sum += util.MustAtoi(c)
		}
	}
	elfs = append(elfs, sum)
	sort.Sort(sort.Reverse(sort.IntSlice(elfs)))
	return elfs
}

func A(input *challenge.Challenge) int {
	return calories(input.Lines())[0]
}

func B(input *challenge.Challenge) int {
	cals := calories(input.Lines())
	return cals[0] + cals[1] + cals[2]
}
