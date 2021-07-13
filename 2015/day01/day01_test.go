package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/1e9y/adventofcode/challenge"
)

var testFloor = []struct {
	input  string
	expect int
}{
	{"(())", 0},
	{"()()", 0},
	{"(((", 3},
	{"(()(()(", 3},
	{"))(((((", 3},
	{"())", -1},
	{"))(", -1},
	{")))", -3},
	{")())())", -3},
}

func TestFloor(t *testing.T) {
	for _, c := range testFloor {
		floor, _ := stairs(c.input)
		assert.Equal(t, c.expect, floor)
	}
}

var testBasement = []struct {
	input  string
	expect int
}{
	{")", 1},
	{"()())", 5},
}

func TestBasement(t *testing.T) {
	for _, c := range testBasement {
		_, basement := stairs(c.input)
		assert.Equal(t, c.expect, basement)
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, A(input), 280)
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, B(input), 1797)
}
