package day02

import (
	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func safe(input []int, tolerance int) bool {
	// t := tolerance
	for i := 1; i < len(input); i++ {
		abs := util.AbsInt(input[i] - input[i-1])
		// fmt.Println(input, input[i], tolerance)

		if ((input[i] < input[i-1]) == (input[1] < input[0])) &&
			(abs >= 1 && 3 >= abs) {
			continue
		}

		tolerance--
		if tolerance < 0 {
			return false
		}
	}

	return true
}

func A(input *challenge.Challenge) (answer int) {
	for _, report := range input.Matrix(challenge.WithSeparator(" ")) {
		if safe(report, 0) {
			answer++
		}
	}

	return
}

func B(input *challenge.Challenge) (answer int) {
	for _, report := range input.Matrix(challenge.WithSeparator(" ")) {
		if safe(report, 1) {
			answer++
		}
	}

	return
}
