package day15

import (
	"math"

	"github.com/1e9y/adventofcode/challenge"
)

var directions = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func path(input [][]int) int {
	height := len(input)
	width := len(input[0])

	distances := make([][]int, height)
	for y := range distances {
		distances[y] = make([]int, width)
		for x := range distances[y] {
			distances[y][x] = math.MaxInt32
		}
	}

	distances[0][0] = 0
	queue := [][]int{{0, 0}}

	var dx, dy, sx, sy, alt int
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sx = node[0]
		sy = node[1]

		for _, d := range directions {
			dx = sx + d[0]
			dy = sy + d[1]

			if dx < 0 || dy < 0 || dx >= width || dy >= height {
				continue
			}

			alt = distances[sy][sx] + input[dy][dx]
			if alt < distances[dy][dx] {
				queue = append(queue, []int{dx, dy})
				distances[dy][dx] = alt
			}
		}
	}

	return distances[height-1][width-1]
}

func multiply(matrix [][]int, multiplier int) [][]int {
	height := len(matrix)
	width := len(matrix[0])

	result := make([][]int, height*multiplier)
	for y := range result {
		result[y] = make([]int, width*multiplier)
		for x := range result[y] {
			value := matrix[y%height][x%width] + x/width + y/height
			if value > 9 {
				value = value%10 + 1
			}
			result[y][x] = value
		}
	}

	return result
}

func A(input *challenge.Challenge) int {
	return path(input.Matrix())
}

func B(input *challenge.Challenge) int {
	return path(multiply(input.Matrix(), 5))
}
