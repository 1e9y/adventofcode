package day01

import (
	"fmt"
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

func tests() []challenge.TestCase[int] {
	return []challenge.TestCase[int]{
		{
			`3   4
4   3
2   5
1   3
3   9
3   3`,
			[]int{11, 31},
		},
	}
}

func TestDistance(t *testing.T) {
	for i, tt := range tests() {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(tt.Input)
			assert.Equal(t, tt.Want[0], A(input))
		})
	}
}

func TestSimilarity(t *testing.T) {
	for _, tt := range tests() {
		t.Run(tt.Name(), func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(tt.Input)
			assert.Equal(t, tt.Want[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 1222801, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 22545250, B(input))
}
