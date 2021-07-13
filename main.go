package main

import (
	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/cmd"
)

type SolutionFunc func(input *challenge.Challenge) int

type Solutions map[int]map[string]SolutionFunc

func (s Solutions) Register(day int, part string, solution SolutionFunc) {
	if _, ok := s[day]; !ok {
		s[day] = make(map[string]SolutionFunc)
	}
	s[day][part] = solution
}

var solutions Solutions

//go:generate go run ./gen
func main() {
	cmd.Execute()
}
