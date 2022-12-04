package day03

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
		`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`,
		[]int{157, 70},
	},
}

func TestPriority(t *testing.T) {
	cases := map[rune]int{
		'a': 1,
		'm': 13,
		'z': 26,
		'A': 27,
		'M': 39,
		'Z': 52,
	}
	for c, p := range cases {
		assert.Equal(t, p, priority(c))
	}
}

func TestRucksack(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:3], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestGroup(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:3], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 7553, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 2758, B(input))
}
