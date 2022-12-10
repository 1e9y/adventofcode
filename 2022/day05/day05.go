package day05

import (
	"regexp"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type stack []rune

func top(s stack) rune {
	if len(s) == 0 {
		return ' '
	}
	return s[len(s)-1]
}

func move(from, to *stack, count int) {
	length := len(*from)
	if length == 0 {
		return
	}
	*to = append(*to, (*from)[length-count:]...)
	*from = (*from)[:length-count]
}

func shift(s *stack, crate rune) {
	*s = append([]rune{crate}, *s...)
}

func (s stack) String() (result string) {
	for _, c := range s {
		result += string(c)
	}
	return
}

func parseStacks(line string, stacks []stack) {
	for i, c := range line {
		if 'A' <= c && c <= 'Z' {
			p := (i - 1) / 4
			shift(&stacks[p], c)
		}
	}
}

var moveRe = regexp.MustCompile(`^move (\d+) from (\d) to (\d)$`)

func parseMove(line string, stacks []stack, once bool) {
	matches := moveRe.FindStringSubmatch(line)
	if len(matches) != 4 {
		return
	}
	count := util.MustAtoi(matches[1])
	from := util.MustAtoi(matches[2]) - 1
	to := util.MustAtoi(matches[3]) - 1
	if once {
		for i := 0; i < count; i++ {
			move(&stacks[from], &stacks[to], 1)
		}
	} else {
		move(&stacks[from], &stacks[to], count)
	}
}

func supplies(input <-chan string, once bool) (result string) {
	var operator func(line string, stacks []stack)
	operator = parseStacks
	stacks := make([]stack, 9)
	for i := range stacks {
		stacks[i] = stack{}
	}

	for line := range input {
		if line == "" {
			operator = func(line string, stacks []stack) {
				parseMove(line, stacks, once)
			}
		}
		operator(line, stacks)
	}

	for i := range stacks {
		result += string(top(stacks[i]))
	}
	return
}

func A(input *challenge.Challenge) string {
	return strings.TrimSpace(supplies(input.Lines(), true))
}

func B(input *challenge.Challenge) string {
	return strings.TrimSpace(supplies(input.Lines(), false))
}
