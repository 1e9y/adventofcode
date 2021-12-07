package day07

import (
	"math"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func alignment(input string, consumption func(n int) int) int {
	parts := strings.Split(input, ",")
	var positions []int
	var distance int
	for _, n := range parts {
		k := util.MustAtoi(n)
		if k > distance {
			distance = k
		}
		positions = append(positions, k)
	}

	minFuel := math.MaxInt32
	for d := 0; d <= distance; d++ {
		fuel := 0
		for _, p := range positions {
			fuel += consumption(util.AbsInt(p - d))
		}
		minFuel = util.MinInt(fuel, minFuel)
	}

	return minFuel
}

func A(input *challenge.Challenge) int {
	return alignment(input.LineSlice()[0], func(n int) int { return n })
}

func B(input *challenge.Challenge) int {
	return alignment(input.LineSlice()[0], func(n int) int { return n * (n + 1) / 2 })
}
