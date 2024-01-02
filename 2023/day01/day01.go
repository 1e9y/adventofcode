package day01

import (
	"fmt"
	"regexp"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

var pattern = regexp.MustCompile(`\d`)

func numeric(input string) int {
	var a, b int
	matches := pattern.FindAllString(input, -1)
	if len(matches) < 1 {
		panic(fmt.Sprintf("bad input: no number in string: %s", input))
	}
	a = util.MustAtoi(matches[0])
	if len(matches) == 1 {
		b = a
	} else {
		b = util.MustAtoi(matches[len(matches)-1])
	}
	return a*10 + b
}

var index = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func alphanumeric(input string) int {
	var a, b int
	var pattern = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	matches := pattern.FindAllString(input, -1)

	if len(matches) < 1 {
		panic(fmt.Sprintf("bad input: no number in string: %s", input))
	}

	if v, ok := index[matches[0]]; ok {
		a = v
	} else {
		a = util.MustAtoi(matches[0])
	}

	if len(matches) == 1 {
		b = a
	} else {
		if v, ok := index[matches[len(matches)-1]]; ok {
			b = v
		} else {
			b = util.MustAtoi(matches[len(matches)-1])
		}
	}

	fmt.Println(input, matches, a, b)

	return a*10 + b
}

func calibration(input <-chan string) (sums []int) {
	sums = make([]int, 2)
	for c := range input {
		// sums[0] += numeric(c)
		sums[1] += alphanumeric(c)
	}
	return
}

func A(input *challenge.Challenge) int {
	return calibration(input.Lines())[0]
}

func B(input *challenge.Challenge) int {
	return calibration(input.Lines())[1]
}
