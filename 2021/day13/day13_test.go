package day13

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  string
	expect []interface{}
}{
	{
		`6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`,
		[]interface{}{17, 16},
	},
}

func TestOneFold(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:12], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestAllFolds(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:12], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 682, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()

	// ####  ##   ##  #  # ###  #### #  # ####
	// #    #  # #  # #  # #  #    # #  # #
	// ###  #  # #    #  # #  #   #  #### ###
	// #    #### # ## #  # ###   #   #  # #
	// #    #  # #  # #  # # #  #    #  # #
	// #    #  #  ###  ##  #  # #### #  # ####

	assert.Equal(t, 104, B(input))
}
