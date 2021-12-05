package day05

import (
	"fmt"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

var grid = make(map[int]int)

func lineFromString(input string) (result [4]int) {
	parts := strings.Split(input, " -> ")
	start := strings.Split(parts[0], ",")
	end := strings.Split(parts[1], ",")
	result[0] = util.MustAtoi(start[0])
	result[1] = util.MustAtoi(start[1])
	result[2] = util.MustAtoi(end[0])
	result[3] = util.MustAtoi(end[1])
	return
}

func overlap(input <-chan string) (result int) {
	for s := range input {
		line := lineFromString(s)

		maxX := util.MaxInt(line[0], line[2])
		minX := util.MinInt(line[0], line[2])
		maxY := util.MaxInt(line[1], line[3])
		minY := util.MinInt(line[1], line[3])

		if line[0] == line[2] || line[1] == line[3] {
			for i := minX; i <= maxX; i++ {
				for j := minY; j <= maxY; j++ {
					grid[i*1000+j]++
					if grid[i*1000+j] == 2 {
						result++
					}
				}
			}
		}
	}
	return
}

func allOverlap(input <-chan string) (result int) {
	for s := range input {
		line := lineFromString(s)

		maxX := util.MaxInt(line[0], line[2])
		minX := util.MinInt(line[0], line[2])
		maxY := util.MaxInt(line[1], line[3])
		minY := util.MinInt(line[1], line[3])

		if line[0] == line[2] || line[1] == line[3] {
			for i := minX; i <= maxX; i++ {
				for j := minY; j <= maxY; j++ {
					grid[i*1000+j]++
					if grid[i*1000+j] == 2 {
						result++
					}
				}
			}
		} else {
			difx := util.AbsInt(line[0] - line[2])
			dify := util.AbsInt(line[1] - line[3])
			if difx != dify {
				panic(fmt.Sprintf("bad input: %v", line))
			}

			sx := line[0]
			sy := line[1]
			ex := line[2]
			ey := line[3]
			var dx, dy int
			if sx < ex {
				dx = 1
			} else {
				dx = -1
			}
			if sy < ey {
				dy = 1
			} else {
				dy = -1
			}

			// fmt.Println("line", line)
			for i := 0; i <= difx; i++ {
				grid[sx*1000+sy]++
				if grid[sx*1000+sy] == 2 {
					result++
				}
				sx += dx
				sy += dy
			}
		}
	}
	return
}

func A(input *challenge.Challenge) int {
	return overlap(input.Lines())
}

func B(input *challenge.Challenge) int {
	return allOverlap(input.Lines())
}
