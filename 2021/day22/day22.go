package day22

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

var stepRe = regexp.MustCompile(`(on|off) x=([\d-]+)..([\d-]+),y=([\d-]+)..([\d-]+),z=([\d-]+)..([\d-]+)`)

type Step struct {
	status                             bool
	xmin, xmax, ymin, ymax, zmin, zmax int
}

func parseStep(input string) Step {
	matches := stepRe.FindStringSubmatch(input)
	if len(matches) != 8 {
		panic(fmt.Errorf("bad input: %v", input))
	}
	return Step{
		status: matches[1] == "on",
		xmin:   util.MustAtoi(matches[2]),
		xmax:   util.MustAtoi(matches[3]),
		ymin:   util.MustAtoi(matches[4]),
		ymax:   util.MustAtoi(matches[5]),
		zmin:   util.MustAtoi(matches[6]),
		zmax:   util.MustAtoi(matches[7]),
	}
}

func parseInput(input <-chan string, filter func(*Step) bool) (steps []Step) {
	steps = []Step{}
	for line := range input {
		step := parseStep(line)
		if !filter(&step) {
			continue
		}
		steps = append(steps, step)
	}
	return
}

func unique(slice []int) (result []int) {
	set := make(map[int]bool)
	for _, n := range slice {
		set[n] = true
	}
	result = make([]int, 0)
	for k := range set {
		result = append(result, k)
	}
	return
}

func reactor(steps []Step) int {
	xaxis := make([]int, 0)
	yaxis := make([]int, 0)
	zaxis := make([]int, 0)

	for _, step := range steps {
		xaxis = append(xaxis, step.xmin, step.xmax+1)
		yaxis = append(yaxis, step.ymin, step.ymax+1)
		zaxis = append(zaxis, step.zmin, step.zmax+1)
	}

	xaxis = unique(xaxis)
	yaxis = unique(yaxis)
	zaxis = unique(zaxis)
	sort.Ints(xaxis)
	sort.Ints(yaxis)
	sort.Ints(zaxis)

	xmap := make(map[int]int)
	for i, x := range xaxis {
		xmap[x] = i
	}

	ymap := make(map[int]int)
	for j, y := range yaxis {
		ymap[y] = j
	}

	zmap := make(map[int]int)
	for k, x := range zaxis {
		zmap[x] = k
	}

	grid := make([][][]bool, len(xaxis))
	for x := 0; x < len(xaxis); x++ {
		grid[x] = make([][]bool, len(yaxis))
		for y := 0; y < len(yaxis); y++ {
			grid[x][y] = make([]bool, len(zaxis))
			for z := 0; z < len(zaxis); z++ {
				grid[x][y][z] = false
			}
		}
	}

	for _, step := range steps {
		xstart := xmap[step.xmin]
		xend := xmap[step.xmax+1] - 1
		ystart := ymap[step.ymin]
		yend := ymap[step.ymax+1] - 1
		zstart := zmap[step.zmin]
		zend := zmap[step.zmax+1] - 1

		for x := xstart; x <= xend; x++ {
			for y := ystart; y <= yend; y++ {
				for z := zstart; z <= zend; z++ {
					grid[x][y][z] = step.status
				}
			}
		}
	}

	enabled := 0
	for x := 0; x < len(xaxis)-1; x++ {
		for y := 0; y < len(yaxis)-1; y++ {
			for z := 0; z < len(zaxis)-1; z++ {
				if grid[x][y][z] {
					enabled += (xaxis[x+1] - xaxis[x]) * (yaxis[y+1] - yaxis[y]) * (zaxis[z+1] - zaxis[z])
				}
			}
		}
	}

	return enabled
}

func A(input *challenge.Challenge) int {
	return reactor(parseInput(input.Lines(), func(step *Step) bool {
		if 50 < step.xmax || step.xmin < -50 ||
			50 < step.ymax || step.ymin < -50 ||
			50 < step.zmax || step.zmin < -50 {
			return false
		}
		return true
	}))
}

func B(input *challenge.Challenge) int {
	return reactor(parseInput(input.Lines(), func(step *Step) bool {
		return true
	}))
}
