package day06

import (
	"github.com/1e9y/adventofcode/challenge"
)

func marker(input string, length int) (index int) {
	memo := make(map[rune]int)
	for i, c := range input {
		if i >= length {
			p := rune(input[i-length])
			memo[p]--
			if memo[p] <= 0 {
				delete(memo, p)
			}
		}
		memo[c]++
		if len(memo) == length {
			return i + 1
		}
	}
	return
}

func A(input *challenge.Challenge) int {
	return marker(<-input.Lines(), 4)
}

func B(input *challenge.Challenge) int {
	return marker(<-input.Lines(), 14)
}
