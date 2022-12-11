package day06

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  string
	expect []int
}{
	{`mjqjpqmgbljsphdztnvjfqwrcgsmlb`, []int{7, 19}},
	{`bvwbjplbgvbhsrlpgdmjqwftvncz`, []int{5, 23}},
	{`nppdvjthqldpwncqszvftbrmjlhg`, []int{6, 23}},
	{`nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`, []int{10, 29}},
	{`zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`, []int{11, 26}},
}

func TestStartOfPacket(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestStartOfMessage(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 1625, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 2250, B(input))
}
