package day25

import (
	"github.com/1e9y/adventofcode/challenge"
)

const (
	Empty    = 0
	Right    = 1
	Down     = 2
	Both     = Right | Down
	NextBoth = (Right << 2) | (Down << 2)
)

var symbols = map[byte]int{
	'.': Empty,
	'>': Right,
	'v': Down,
}

func move(sea [][]int) bool {
	height := len(sea)
	width := len(sea[0])

	moved := false

	for j := 0; j < height; j++ {
		for i := width - 1; i >= 0; i-- {
			next := i + 1
			if i == width-1 {
				next = 0
			}
			if sea[j][i]&Right != 0 && sea[j][next]&Both == 0 {
				sea[j][next] = sea[j][next] &^ NextBoth
				sea[j][next] = sea[j][next] | (Right << 2)
				sea[j][i] = sea[j][i] &^ NextBoth
				moved = true
			}
		}
	}

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			sea[j][i] = sea[j][i] &^ Both
			sea[j][i] = sea[j][i] | (sea[j][i] >> 2)
		}
	}

	for j := height - 1; j >= 0; j-- {
		for i := 0; i < width; i++ {
			next := j + 1
			if j == height-1 {
				next = 0
			}
			if sea[j][i]&Down != 0 && sea[next][i]&Both == 0 {
				sea[next][i] = sea[next][i] &^ NextBoth
				sea[next][i] = sea[next][i] | (Down << 2)
				sea[j][i] = sea[j][i] &^ NextBoth
				moved = true
			}
		}
	}

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			sea[j][i] = sea[j][i] &^ Both
			sea[j][i] = sea[j][i] | (sea[j][i] >> 2)
		}
	}

	return moved
}

func simulate(sea [][]int) int {
	steps := 0
	for move(sea) {
		steps++
	}
	return steps + 1
}

func parseInput(input []string) [][]int {
	var width, height int
	height = len(input)
	width = len(input[0])

	result := make([][]int, height)
	for i := 0; i < height; i++ {
		result[i] = make([]int, width)
	}

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			result[j][i] = symbols[input[j][i]] | (symbols[input[j][i]] << 2)
		}
	}

	return result
}

func A(input *challenge.Challenge) int {
	return simulate(parseInput(input.LineSlice()))
}

func B(input *challenge.Challenge) int {
	return 2021
}
