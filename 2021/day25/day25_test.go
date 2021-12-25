package day25

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
		input: `v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>`,
		expect: []int{58, 44169},
	},
}

func TestTraffic(t *testing.T) {
	// t.Skip()
	for _, c := range testCases {
		t.Run(c.input[:12], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestUnfoldedBurrow(t *testing.T) {
	t.Skip()
	for _, c := range testCases {
		t.Run(c.input[:12], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	// t.Skip()
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 424, A(input))
}

func TestB(t *testing.T) {
	t.Skip()
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 48759, B(input))
}
