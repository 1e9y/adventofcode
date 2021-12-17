package day16

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

var testCasesA = []struct {
	input  string
	expect int
}{
	{"D2FE28", 6},
	{"38006F45291200", 9},
	{"EE00D40C823060", 14},
	{"8A004A801A8002F478", 16},
	{"620080001611562C8802118E34", 12},
	{"C0015000016115A2E0802F182340", 23},
	{"A0016C880162017C3686B18A3D4780", 31},
	{"C200B40A82", 14},
}

var testCasesB = []struct {
	input  string
	expect int
}{
	{"C200B40A82", 3},
	{"04005AC33890", 54},
	{"880086C3E88112", 7},
	{"CE00C43D881120", 9},
	{"D8005AC2A8F0", 1},
	{"F600BC2D8F", 0},
	{"9C005AC2F8F0", 0},
	{"9C0141080250320F1802104A08", 1},
}

func TestSum(t *testing.T) {
	for _, c := range testCasesA {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect, A(input))
		})
	}
}

func TestResult(t *testing.T) {
	for _, c := range testCasesB {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect, B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 934, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 912901337844, B(input))
}
