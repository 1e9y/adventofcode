package day03

import (
	"fmt"
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/require"
)

var testCases = []struct {
	input  string
	expect []int
}{
	{`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`,
		[]int{4361, 467835}},
}

func TestSchematic(t *testing.T) {
	for i, c := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			require.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestGears(t *testing.T) {
	for i, c := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			require.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	require.Equal(t, 521601, A(input))
}
func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	require.Equal(t, 80694070, B(input))
}
