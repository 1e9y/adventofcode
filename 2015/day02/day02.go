package day02

import (
	"strings"

	"github.com/1e9y/adventofcode/util"

	"github.com/1e9y/adventofcode/challenge"
)

func wrapping(l, w, h int) int {
	return 2*(l*w+w*h+h*l) + util.MinInt(l*w, w*h, h*l)
}

func ribbon(l, w, h int) int {
	return l*w*h + util.MinInt(2*(l+w), 2*(w+h), 2*(h+l))
}

func parse(s string) (int, int, int) {
	d := strings.Split(s, "x")
	return util.MustAtoi(d[0]), util.MustAtoi(d[1]), util.MustAtoi(d[2])
}

func A(input *challenge.Challenge) int {
	result := 0
	for line := range input.Lines() {
		l, w, h := parse(line)
		result += wrapping(l, w, h)
	}
	return result
}

func B(input *challenge.Challenge) int {
	result := 0
	for line := range input.Lines() {
		l, w, h := parse(line)
		result += ribbon(l, w, h)
	}
	return result
}
