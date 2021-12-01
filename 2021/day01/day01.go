package day01

import (
	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func deep(input []string, width int) int {
	var a, b, answer int
	for i := range input {
		if i < width-1 {
			continue
		}

		b = func() int {
			sum := 0
			for n := 0; n < width; n++ {
				sum += util.MustAtoi(input[i-n])
			}
			return sum
		}()

		if b > a {
			answer++
		}
		a = b
	}
	return answer - 1
}

func A(input *challenge.Challenge) int {
	return deep(input.LineSlice(), 1)
}

func B(input *challenge.Challenge) int {
	return deep(input.LineSlice(), 3)
}
