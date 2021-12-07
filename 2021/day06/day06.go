package day06

import (
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func population(input string, period int) int {
	var lanternfish []int
	parts := strings.Split(input, ",")
	for _, n := range parts {
		lanternfish = append(lanternfish, util.MustAtoi(n))
	}

	pop := make([]int, 9)
	for _, n := range lanternfish {
		pop[n]++
	}

	for i := 0; i < period; i++ {
		fish := pop[0]
		pop = pop[1:]
		pop = append(pop, fish)
		pop[6] += fish
	}

	result := 0
	for _, n := range pop {
		result += n
	}
	return result
}

func A(input *challenge.Challenge) int {
	return population(input.LineSlice()[0], 80)
}

func B(input *challenge.Challenge) int {
	return population(input.LineSlice()[0], 256)
}
