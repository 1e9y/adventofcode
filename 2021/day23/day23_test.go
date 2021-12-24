package day23

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
		input: `#############
#...........#
###B#C#B#D###
  #A#D#C#A#  
  #########  `, // here trailing spaces matter!
		expect: []int{12521, 44169},
	},
}

func TestBurrow(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[30:39], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestUnfoldedBurrow(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[30:39], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 14467, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 48759, B(input))
}
