package day20

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  string
	expect []int
}{
	{
		`..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`,
		[]int{35, 3351},
	},
}

func TestEnhancement(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:12], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestBetterEnhancement(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:12], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 5326, A(input))
}

func TestB(t *testing.T) {
	t.Skip()
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 17096, B(input))
}
