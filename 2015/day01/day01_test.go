package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/1e9y/adventofcode/challenge"
)

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile("")
	assert.Equal(t, A(input), 280)
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile("")
	assert.Equal(t, B(input), 280)
}
