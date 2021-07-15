package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/1e9y/adventofcode/challenge"
)

var testCases = []struct {
	input  string
	expect int
}{
	{"abcdef", 609043},
	{"pqrstuv", 1048970},
}

func TestMine(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect, A(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 346386, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 9958218, B(input))
}
