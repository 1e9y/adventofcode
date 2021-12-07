package day07

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
		"16,1,2,0,4,2,7,1,2,14",
		[]int{37, 168},
	},
}

func TestAlignment(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestExpensiveAlignment(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 336131, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 92676646, B(input))
}
