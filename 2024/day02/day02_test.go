package day02

import (
	"fmt"
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

func tests() []challenge.TestCase[int] {
	return []challenge.TestCase[int]{
		{
			`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`,
			[]int{2, 4},
		},
	}
}

func TestSafe(t *testing.T) {
	for i, tt := range tests() {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(tt.Input)
			assert.Equal(t, tt.Want[0], A(input))
		})
	}
}

func TestDampener(t *testing.T) {
	for _, tt := range tests() {
		t.Run(tt.Name(), func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(tt.Input)
			assert.Equal(t, tt.Want[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 639, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 674, B(input))
}
