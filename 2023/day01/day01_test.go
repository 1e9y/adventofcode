package day01

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  string
	expect []int
}{
	// 	{`1abc2
	// pqr3stu8vwx
	// a1b2c3d4e5f
	// treb7uchet
	// `,
	// 		[]int{142, 142},
	// },
	{`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
`,
		[]int{281, 281},
	},
}

func TestCalibration(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:4], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestTrueCalibration(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:4], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 56465, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, -1, B(input))
}
