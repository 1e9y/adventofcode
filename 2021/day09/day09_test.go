package day09

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
		`2199943210
3987894921
9856789892
8767896789
9899965678`,
		[]int{15, 1134},
	},
}

func TestLowPoints(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:10], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestBasins(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:10], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 439, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 900900, B(input))
}
