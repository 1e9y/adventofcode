package day02

import (
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func navigation(input <-chan string) int {
	var position, depth int
	for s := range input {
		parts := strings.Split(s, " ")
		switch parts[0] {
		case "forward":
			position += util.MustAtoi(parts[1])
		case "down":
			depth += util.MustAtoi(parts[1])
		case "up":
			depth -= util.MustAtoi(parts[1])
		}
	}
	return position * depth
}

func aim(input <-chan string) int {
	var position, depth, aim int
	for s := range input {
		parts := strings.Split(s, " ")
		switch parts[0] {
		case "forward":
			position += util.MustAtoi(parts[1])
			depth += aim * util.MustAtoi(parts[1])
		case "down":
			aim += util.MustAtoi(parts[1])
		case "up":
			aim -= util.MustAtoi(parts[1])
		}
	}
	return position * depth
}

func A(input *challenge.Challenge) int {
	return navigation(input.Lines())
}

func B(input *challenge.Challenge) int {
	return aim(input.Lines())
}
