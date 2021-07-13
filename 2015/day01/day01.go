package day01

import (
	"github.com/1e9y/adventofcode/challenge"
)

func stairs(input string) (int, int) {
	var f, b int
	for i, c := range input {
		if c == '(' {
			f++
		} else {
			f--
		}
		if b == 0 && f == -1 {
			b = i + 1
		}
	}
	return f, b
}

func A(input *challenge.Challenge) int {
	var answer int
	answer, _ = stairs(input.LineSlice()[0])
	return answer
}

func B(input *challenge.Challenge) int {
	var answer int
	_, answer = stairs(input.LineSlice()[0])
	return answer
}
