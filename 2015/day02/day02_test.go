package day02

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  string
	expect []int
}{
	{"2x3x4", []int{58, 34}},
	{"1x1x10", []int{43, 14}},
}

func TestWrappings(t *testing.T) {
	for _, c := range testCases {
		w := wrapping(parse(c.input))
		r := ribbon(parse(c.input))
		assert.Equal(t, c.expect[1], r)
		assert.Equal(t, c.expect[0], w)
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, A(input), 1598415)
}
func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, B(input), 3812909)
}
