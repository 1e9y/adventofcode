package day03

import (
	"fmt"
	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func priority(char rune) int {
	if char >= 'a' && char <= 'z' {
		return int(char-'a') + 1
	}
	if char >= 'A' && char <= 'Z' {
		return int(char-'A') + 27
	}
	panic(fmt.Errorf("bad priority: %v", char))
}

func share(content ...string) int {
	index := make([]uint64, len(content))
	for i, c := range content {
		for j := 0; j < len(c); j++ {
			index[i] |= 1 << uint64(priority(rune(c[j])))
		}
	}
	var badge uint64
	for k, x := range index {
		if k == 0 {
			badge = x
		} else {
			badge &= x
		}
	}
	return util.RightmostBit(int(badge))
}

func rucksack(input <-chan string) int {
	sum := 0
	for line := range input {
		m := len(line) / 2
		s := share(line[:m], line[m:])
		sum += s
	}
	return sum
}

func groups(input <-chan string) int {
	sum := 0
	for {
		a, ok := <-input
		if !ok {
			break
		}
		b, ok := <-input
		if !ok {
			break
		}
		c, ok := <-input
		if !ok {
			break
		}
		s := share(a, b, c)
		sum += s
	}
	return sum
}

func A(input *challenge.Challenge) int {
	return rucksack(input.Lines())
}

func B(input *challenge.Challenge) int {
	return groups(input.Lines())
}
