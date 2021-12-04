package day03

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  string
	expect []int
}{
	{`00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`,
		[]int{198, 230},
	},
}

func TestConsumption(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestRating(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 4006064, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 5941884, B(input))
}
