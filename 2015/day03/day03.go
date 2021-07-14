package day03

import (
	"fmt"

	"github.com/1e9y/adventofcode/challenge"
)

func key(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}

var moves = map[rune][]int{
	'^': {0, -1},
	'v': {0, 1},
	'>': {1, 0},
	'<': {-1, 0},
}

func gifted(mem map[string]bool) int {
	total := 0
	for _, b := range mem {
		if b {
			total++
		}
	}
	return total
}

type Santa struct {
	mem  *map[string]bool
	x, y int
}

func (s *Santa) move(m rune) {
	d := moves[m]
	s.x += d[0]
	s.y += d[1]
}

func (s *Santa) visited() bool {
	k := key(s.x, s.y)
	return (*s.mem)[k]
}

func (s *Santa) give() {
	k := key(s.x, s.y)
	(*s.mem)[k] = true
}

func newSanta(mem *map[string]bool) *Santa {
	s := &Santa{
		x:   0,
		y:   0,
		mem: mem,
	}
	s.give()
	return s
}

func A(input *challenge.Challenge) int {
	mem := make(map[string]bool)
	santa := newSanta(&mem)
	for line := range input.Lines() {
		for _, m := range line {
			santa.move(m)
			if !santa.visited() {
				santa.give()
			}
		}
	}
	return gifted(mem)
}

func B(input *challenge.Challenge) int {
	mem := make(map[string]bool)
	santa := newSanta(&mem)
	robot := newSanta(&mem)
	actor := &santa
	for line := range input.Lines() {
		for _, m := range line {
			(*actor).move(m)
			if !(*actor).visited() {
				(*actor).give()
			}

			if actor == &santa {
				actor = &robot
			} else {
				actor = &santa
			}
		}
	}
	return gifted(mem)
}
