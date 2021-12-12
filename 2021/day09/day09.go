package day09

import (
	"sort"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func adjacent(x, y int, input [][]int) (result [][]int) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			sx := x + dx
			sy := y + dy
			if sx >= 0 && sy >= 0 &&
				sx < len(input[0]) && sy < len(input) &&
				util.AbsInt(dx) != util.AbsInt(dy) {
				result = append(result, []int{sx, sy})
			}
		}
	}
	return
}

func basins(input [][]int) (int, int) {
	visited := make(map[string]bool)

	sizes := make([]int, 0)
	bottoms := 0

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if visited[util.Key(x, y)] || input[y][x] == 9 {
				continue
			}
			queue := [][]int{{x, y}}
			size := 0
			bottom := 9
			for len(queue) > 0 {
				p := queue[0]
				queue = queue[1:]
				x := p[0]
				y := p[1]
				key := util.Key(x, y)
				if visited[key] {
					continue
				}
				visited[key] = true
				size++
				if input[y][x] < bottom {
					bottom = input[y][x]
				}

				for _, a := range adjacent(x, y, input) {
					ax := a[0]
					ay := a[1]
					if !visited[util.Key(ax, ay)] && input[ay][ax] < 9 {
						queue = append(queue, []int{ax, ay})
					}
				}
			}
			bottoms += bottom + 1
			sizes = append(sizes, size)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	return bottoms, sizes[0] * sizes[1] * sizes[2]
}

func A(input *challenge.Challenge) int {
	answer, _ := basins(input.Matrix())
	return answer
}

func B(input *challenge.Challenge) int {
	_, answer := basins(input.Matrix())
	return answer
}
