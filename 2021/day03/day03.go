package day03

import (
	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

func consumption(input <-chan string) int {
	var gamma, epsilon, size int
	index := make(map[int]int)

	for line := range input {
		for i, c := range line {
			if c == '1' {
				index[i]++
			}
		}
		size++
	}

	for i := 0; i < len(index); i++ {
		if index[i] < size/2 {
			gamma |= 1
			gamma <<= 1
			epsilon <<= 1
		} else {
			epsilon |= 1
			gamma <<= 1
			epsilon <<= 1
		}
	}
	gamma >>= 1
	epsilon >>= 1

	return gamma * epsilon
}

type Criteria func(population int) byte

type Report []string

func (r Report) reduce(pos int, criteria Criteria) (result []string) {
	index := make(map[int]int)
	for _, line := range r {
		for i, c := range line {
			if c == '0' {
				index[i] -= 1
			} else {
				index[i] += 1
			}
		}
	}
	for _, line := range r {
		if line[pos] == criteria(index[pos]) {
			result = append(result, line)
		}
	}
	return
}

func (d Report) filterBy(criteria Criteria) string {
	var pos int
	for len(d) > 1 {
		d = d.reduce(pos, criteria)
		pos++
	}
	return d[0]
}

func rating(input []string) int {
	report := Report(input)

	oxygen := report.filterBy(func(population int) byte {
		if population >= 0 {
			return '1'
		}
		return '0'
	})

	carbondioxide := report.filterBy(func(population int) byte {
		if population >= 0 {
			return '0'
		}
		return '1'
	})

	return util.MustParseInt(oxygen, 2) * util.MustParseInt(carbondioxide, 2)
}

func A(input *challenge.Challenge) int {
	return consumption(input.Lines())
}

func B(input *challenge.Challenge) int {
	return rating(input.LineSlice())
}
