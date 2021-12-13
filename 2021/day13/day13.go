package day13

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type Instruction struct {
	axis  rune
	value int
}

func render(coordinates [][]int) (string, int) {
	var width, height int
	for _, c := range coordinates {
		if c[0] > width {
			width = c[0]
		}
		if c[1] > height {
			height = c[1]
		}

	}

	height++
	width++

	index := make([][]bool, height)
	for i := range index {
		index[i] = make([]bool, width)
	}

	for _, c := range coordinates {
		index[c[1]][c[0]] = true
	}

	pattern := strings.Builder{}
	count := 0
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if index[j][i] {
				count++
				pattern.WriteString("#")
			} else {
				pattern.WriteString(" ")
			}
		}
		pattern.WriteString("\n")
	}

	return pattern.String(), count
}

func parseInput(input <-chan string) (coordinates [][]int, instructions []Instruction) {
	coordinates = make([][]int, 0)
	instructions = make([]Instruction, 0)

	var parse func(string)

	parseCoordinate := func(line string) {
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			panic(fmt.Errorf("bad input: %v", line))
		}
		coordinates = append(coordinates, []int{
			util.MustAtoi(parts[1]), util.MustAtoi(parts[0]),
		})
	}

	parseInstruction := func(line string) {
		re := regexp.MustCompile(`([xy])=(\d+)$`)
		matches := re.FindStringSubmatch(line)
		if len(matches) != 3 {
			panic(fmt.Errorf("bad input: %v", line))
		}
		instructions = append(instructions, Instruction{
			axis:  rune(matches[1][0]),
			value: util.MustAtoi(matches[2]),
		})
	}

	parse = parseCoordinate

	for line := range input {
		if line == "" {
			parse = parseInstruction
			continue
		}
		parse(line)
	}

	return
}

func fold(input <-chan string, n int) (string, int) {
	coordinates, instructions := parseInput(input)

	for i, instruction := range instructions {
		if i >= n {
			break
		}

		axis := instruction.axis
		value := instruction.value

		for _, c := range coordinates {
			i := 'y' - axis
			if c[i] > value {
				c[i] = value*2 - c[i]
			}
		}
	}

	return render(coordinates)
}

func A(input *challenge.Challenge) int {
	_, count := fold(input.Lines(), 1)
	return count
}

func B(input *challenge.Challenge) int {
	_, count := fold(input.Lines(), math.MaxInt32)
	return count
}
