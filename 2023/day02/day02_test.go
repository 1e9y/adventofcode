package day02

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
	{`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`,
		[]int{8, 2286}},
}

func TestGames(t *testing.T) {
	for i, c := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			require.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestPower(t *testing.T) {
	for i, c := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			require.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	require.Equal(t, 2101, A(input))
}
func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	require.Equal(t, 58269, B(input))
}
