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
	{
		`A Y
B X
C Z`,
		[]int{15, 12},
	},
}

func TestStrategy(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:3], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestRealStrategy(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:3], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 12156, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 10835, B(input))
}
