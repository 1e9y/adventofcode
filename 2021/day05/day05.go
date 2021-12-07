package day05

import (
	"fmt"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type Line struct {
	x1, y1, x2, y2         int
	minX, minY, maxX, maxY int
	len                    int
}

func newLineFromString(input string) Line {
	parts := strings.Split(input, " -> ")
	if len(parts) != 2 {
		panic(fmt.Sprintf("bad input: %s", input))
	}
	start := strings.Split(parts[0], ",")
	end := strings.Split(parts[1], ",")
	if len(start) != 2 || len(end) != 2 {
		panic(fmt.Sprintf("bad input: %s", input))
	}

	line := Line{
		x1: util.MustAtoi(start[0]),
		y1: util.MustAtoi(start[1]),
		x2: util.MustAtoi(end[0]),
		y2: util.MustAtoi(end[1]),
	}

	dx := util.AbsInt(line.x2 - line.x1)
	dy := util.AbsInt(line.y2 - line.y1)
	if (line.x1 != line.x2 && line.y1 != line.y2) && dx != dy {
		panic(fmt.Sprintf("bad input: uneven slopes: %s", input))
	}
	line.len = dx

	line.minX = util.MinInt(line.x1, line.x2)
	line.minY = util.MinInt(line.y1, line.y2)
	line.maxX = util.MaxInt(line.x1, line.x2)
	line.maxY = util.MaxInt(line.y1, line.y2)

	return line
}

func xytoi(x, y int) int {
	return 1000*x + y
}

func overlap(input <-chan string) (a, b int) {
	var grid = make(map[int]*[2]int)

	for s := range input {
		line := newLineFromString(s)

		if line.x1 == line.x2 || line.y1 == line.y2 {
			for i := line.minX; i <= line.maxX; i++ {
				for j := line.minY; j <= line.maxY; j++ {
					key := xytoi(i, j)
					if grid[key] == nil {
						grid[key] = new([2]int)
					}

					grid[key][0]++
					if grid[key][0] == 2 {
						a++
					}
					grid[key][1]++
					if grid[key][1] == 2 {
						b++
					}
				}
			}
		} else {
			var dx, dy int
			if line.x1 < line.x2 {
				dx = 1
			} else {
				dx = -1
			}
			if line.y1 < line.y2 {
				dy = 1
			} else {
				dy = -1
			}

			sx := line.x1
			sy := line.y1
			for i := 0; i <= line.len; i++ {
				key := xytoi(sx, sy)
				if grid[key] == nil {
					grid[key] = new([2]int)
				}

				grid[key][1]++
				if grid[key][1] == 2 {
					b++
				}
				sx += dx
				sy += dy
			}
		}
	}
	return
}

func A(input *challenge.Challenge) int {
	answer, _ := overlap(input.Lines())
	return answer
}

func B(input *challenge.Challenge) int {
	_, answer := overlap(input.Lines())
	return answer
}
