package util

import (
	"fmt"
)

func PrintMatrixln(input [][]int) {
	max := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			max = MaxInt(max, input[y][x])
		}
	}
	width := 1
	for max > 0 {
		max /= 10
		width++
	}

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			fmt.Printf("%*d", width, input[y][x])
		}
		fmt.Println()
	}
	fmt.Println()
}
