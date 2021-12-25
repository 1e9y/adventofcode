package day25

import (
	"fmt"

	"github.com/1e9y/adventofcode/challenge"
)

const (
	Empty = 0
	Right = 1
	Down  = 2
)

var symbolmap = map[int]string{
	0: ".",
	1: ">",
	2: "v",
}

func render(sea [][][]rune) {
	for j := 0; j < len(sea); j++ {
		for i := 0; i < len(sea[0]); i++ {
			fmt.Printf("%c", sea[j][i][0])
		}
		fmt.Println()
	}
	fmt.Println()
}

func move(sea [][][]rune) bool {
	height := len(sea)
	width := len(sea[0])
	moved := false

	for j := 0; j < height; j++ {
		for i := width - 1; i >= 0; i-- {
			next := i + 1
			if i == width-1 {
				next = 0
			}
			if sea[j][i][0] == '>' && sea[j][next][0] == '.' {
				sea[j][next][1] = '>'
				sea[j][i][1] = '.'
				moved = true
			}
		}
	}

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			sea[j][i][0] = sea[j][i][1]
		}
	}

	for j := height - 1; j >= 0; j-- {
		for i := 0; i < width; i++ {
			next := j + 1
			if j == height-1 {
				next = 0
			}
			if sea[j][i][0] == 'v' && sea[next][i][0] == '.' {
				sea[next][i][1] = 'v'
				sea[j][i][1] = '.'
				moved = true
			}
		}
	}

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			sea[j][i][0] = sea[j][i][1]
		}
	}

	return moved
}

func simulate(sea [][][]rune) int {
	// render(sea)
	// for i := 0; i < 3; i++ {
	// 	move(sea)
	// 	println()
	// 	render(sea)
	// }
	steps := 0
	for move(sea) {
		steps++
	}
	// println(steps + 1)
	// println(Right << 2 & 0)
	return steps + 1
}

func parseInput(input []string) [][][]rune {
	var width, height int
	height = len(input)
	width = len(input[0])

	result := make([][][]rune, height)
	for i := 0; i < height; i++ {
		result[i] = make([][]rune, width)
	}

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			result[j][i] = []rune{rune(input[j][i]), rune(input[j][i])}
		}
	}

	return result
}

func A(input *challenge.Challenge) int {
	return simulate(parseInput(input.LineSlice()))
}

func B(input *challenge.Challenge) int {
	return 1
}
