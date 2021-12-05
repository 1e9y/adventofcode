package day05

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  string
	expect []int
}{
	{`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`,
		[]int{5, 12},
	},
}

func TestOverlaps(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:10], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}
func TestAllOverlaps(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:10], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 6113, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 20373, B(input))
}
