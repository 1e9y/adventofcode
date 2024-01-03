package day01

import (
	"fmt"
	"regexp"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

var patternNumeric = regexp.MustCompile(`\d`)

func numeric(input string) int {
	var a, b int
	matches := patternNumeric.FindAllString(input, -1)
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

var patternAlphanumeric = regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
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

	match := patternAlphanumeric.FindString(input)
	if match == "" {
		panic(fmt.Sprintf("bad input: no number in string: %s", input))
	}

	if v, ok := index[match]; ok {
		a = v
	} else {
		a = util.MustAtoi(match)
	}

	var matchb string
	for i := len(input) - 1; i >= 0; i-- {
		matchb = patternAlphanumeric.FindString(input[i:])
		if matchb == "" {
			continue
		}
		if v, ok := index[matchb]; ok {
			b = v
		} else {
			b = util.MustAtoi(matchb)
		}
		break
	}

	return a*10 + b
}

func calibration(input <-chan string, value func(string) int) (sum int) {
	for c := range input {
		sum += value(c)
	}
	return
}

func A(input *challenge.Challenge) int {
	return calibration(input.Lines(), numeric)
}

func B(input *challenge.Challenge) int {
	return calibration(input.Lines(), alphanumeric)
}
