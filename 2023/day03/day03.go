package day03

import (
	"regexp"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func symbol(r byte) bool {
	if r == '.' || (r >= '0' && r <= '9') {
		return false
	}
	return true
}

func part(line int, indices []int, input []string) bool {
	for x := indices[0] - 1; x <= indices[1]; x++ {
		if x < 0 || x >= len(input[0]) {
			continue
		}
		for y := line - 1; y <= line+1; y++ {
			if y < 0 || y >= len(input) {
				continue
			}
			if symbol(input[y][x]) {
				return true
			}
		}
	}
	return false
}

var numberRe = regexp.MustCompile(`\d+`)

func schematic(input []string) (sum int) {
	for i, row := range input {
		indices := numberRe.FindAllStringIndex(row, -1)
		for _, p := range indices {
			if part(i, p, input) {
				sum += util.MustAtoi(row[p[0]:p[1]])
			}
		}
	}
	return
}

func walk(line int, indices []int, input []string, visit func(rune, int, int)) {
	for x := indices[0] - 1; x <= indices[1]; x++ {
		if x < 0 || x >= len(input[0]) {
			continue
		}
		for y := line - 1; y <= line+1; y++ {
			if y < 0 || y >= len(input) {
				continue
			}
			visit(rune(input[y][x]), x, y)
		}
	}
}

func extract(line int, indices []int, input []string) (result int) {

	for i := indices[0]; i < indices[1]; i++ {
		result = result*10 + (int(input[line][i]) - '0')
	}
	return
}

func gears(input []string) (result int) {
	index := make(map[int][]int)

	for i, row := range input {
		indices := numberRe.FindAllStringIndex(row, -1)
		for _, p := range indices {
			walk(i, p, input, func(r rune, i1, i2 int) {
				if r == '*' {
					idx := i2*len(input) + i1
					if _, ok := index[idx]; !ok {
						index[idx] = make([]int, 0, 2)
					}
					index[idx] = append(index[idx], extract(i, p, input))
				}
			})

		}
	}

	for _, g := range index {
		if len(g) == 2 {
			result += g[0] * g[1]
		}
	}

	return
}

func A(input *challenge.Challenge) int {
	return schematic(input.LineSlice())
}

func B(input *challenge.Challenge) int {
	return gears(input.LineSlice())
}
