package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/1e9y/adventofcode/challenge"
)

var testCases = []struct {
	input  string
	expect []int
}{
	{">", []int{2, 2}},
	{"^>", []int{3, 3}},
	{"^>v<", []int{4, 3}},
	{"^v^v^v^v^v", []int{2, 11}},
}

func TestSanta(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestRobot(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 2592, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 2360, B(input))
}
