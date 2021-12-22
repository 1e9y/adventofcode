package day21

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
		`Player 1 starting position: 4
Player 2 starting position: 8`,
		[]int{739785, 444356092776315},
	},
}

func TestDeterministic(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestDirac(t *testing.T) {
	t.Skip()
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 556206, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 630797200227453, B(input))
}
