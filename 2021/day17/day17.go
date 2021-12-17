package day17

import (
	"fmt"
	"regexp"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type Trench struct {
	xmin, xmax, ymin, ymax int
}

func newTrenchFromInput(input string) Trench {
	matches := regexp.MustCompile(`x=(-?[\d]+)..(-?[\d]+), y=(-?[\d]+)..(-?[\d]+)$`).FindStringSubmatch(input)
	if len(matches) != 5 {
		panic(fmt.Errorf("bad input: %s", input))
	}
	return Trench{
		xmin: util.MustAtoi(matches[1]),
		xmax: util.MustAtoi(matches[2]),
		ymin: util.MustAtoi(matches[3]),
		ymax: util.MustAtoi(matches[4]),
	}
}

func heighest(trench Trench) int {
	return util.AbsInt(trench.ymin) * (util.AbsInt(trench.ymin) - 1) / 2
}

func simulateX(vx int, trench Trench) bool {
	x := 0
	for x <= trench.xmax && vx != 0 {
		x += vx
		if vx > 0 {
			vx = util.MaxInt(0, vx-1)
		} else {
			vx = util.MinInt(0, vx+1)
		}
		if x >= trench.xmin && x <= trench.xmax {
			return true
		}
	}
	return false
}

func simulateY(vy int, trench Trench) bool {
	y := 0
	for y >= trench.ymin {
		y += vy
		vy--
		if y >= trench.ymin && y <= trench.ymax {
			return true
		}
	}
	return false
}

func simulate(velocity []int, trench Trench) bool {
	var x, y int
	vx := velocity[0]
	vy := velocity[1]
	for x <= trench.xmax && y >= trench.ymin {
		x += vx
		y += vy
		if vx > 0 {
			vx = util.MaxInt(0, vx-1)
		} else {
			vx = util.MinInt(0, vx+1)
		}
		vy--
		if x >= trench.xmin && x <= trench.xmax && y >= trench.ymin && y <= trench.ymax {
			return true
		}
	}
	return false
}

func velocities(trench Trench) (count int) {
	xx := []int{}
	for x := 0; x <= trench.xmax; x++ {
		if simulateX(x, trench) {
			xx = append(xx, x)
		}
	}

	yy := []int{}
	ymax := heighest(trench)
	for y := trench.ymin; y <= ymax; y++ {
		if simulateY(y, trench) {
			yy = append(yy, y)
		}
	}

	for _, x := range xx {
		for _, y := range yy {
			if simulate([]int{x, y}, trench) {
				count++
			}
		}
	}

	return
}

func A(input *challenge.Challenge) int {
	return heighest(newTrenchFromInput(<-input.Lines()))
}

func B(input *challenge.Challenge) int {
	return velocities(newTrenchFromInput(<-input.Lines()))
}
