// I stopped looking for my own solution after I found brilliant and concious gist by jkseppan:
// https://gist.github.com/jkseppan/1e36172ad4f924a8f86a920e4b1dc1b1
// This is my Go implementation of jkseppan's solution.

package day24

import (
	"fmt"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func parseInput(input []string) [][]int {
	// find lenght of single block
	lenght := 0
	for i, line := range input[1:] {
		if strings.HasPrefix(line, "inp") {
			lenght = i + 1
			break
		}
	}

	// total number of blocks
	blocks := len(input) / lenght

	variables := [][]string{}
outer:
	for l := 1; l < lenght; l++ {
		params := make([]string, blocks)
		for b := 0; b < blocks; b++ {
			parts := strings.Split(input[b*lenght+l], " ")
			if len(parts) != 2 && len(parts) != 3 {
				panic(fmt.Errorf("bad input: %v", input[b*lenght+l]))
			}
			params[b] = parts[len(parts)-1]
		}
		for i := 1; i < len(params); i++ {
			if params[i] != params[i-1] {
				variables = append(variables, params)
				continue outer
			}
		}
	}

	result := make([][]int, len(variables))
	for i := range variables {
		result[i] = make([]int, len(variables[i]))
		for j := range variables[i] {
			result[i][j] = util.MustAtoi(variables[i][j])
		}
	}
	return result
}

func join(ints []int) (result int) {
	for i, n := range ints {
		if i != 0 {
			result *= 10
		}
		result += n
	}
	return
}

func zip(slices ...[]int) [][]int {
	result := make([][]int, len(slices[0]))
	for i := range slices[0] {
		result[i] = []int{}
		for s := range slices {
			result[i] = append(result[i], slices[s][i])
		}
	}
	return result
}

func block(a, b, c, zp, w int) (zz []int) {
	zz = []int{}
	x := zp - w - c
	if x%26 == 0 {
		zz = append(zz, x/26*a)
	}
	if w-b >= 0 && w-b < 26 {
		zz = append(zz, w-b+zp*a)
	}

	return
}

func compute(variables [][]int, ww []int) int {
	aa, bb, cc := variables[0], variables[1], variables[2]
	tuples := zip(aa, bb, cc)

	zz := map[int]int{0: 0}
	result := map[int][]int{}
	var a, b, c int
	for i := len(tuples) - 1; i >= 0; i-- {
		a, b, c = tuples[i][0], tuples[i][1], tuples[i][2]
		nextzz := map[int]int{}
		for _, w := range ww {
			for _, z := range zz {
				for _, z0 := range block(a, b, c, z, w) {
					if _, ok := nextzz[z0]; !ok {
						nextzz[z0] = z0

						if _, ok := result[z]; ok {
							result[z0] = append([]int{w}, result[z]...)
						} else {
							result[z0] = []int{w}
						}
					}
				}

			}
		}
		zz = nextzz
	}

	return join(result[0])
}

func A(input *challenge.Challenge) int {
	return compute(parseInput(input.LineSlice()), []int{9, 8, 7, 6, 5, 4, 3, 2, 1})
}

func B(input *challenge.Challenge) int {
	return compute(parseInput(input.LineSlice()), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}
