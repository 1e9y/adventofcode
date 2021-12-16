package day16

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

var testCasesA = []struct {
	input  string
	expect []int
}{
	{
		"D2FE28",
		[]int{6, 0},
	},
	{
		"38006F45291200",
		[]int{9, 0},
	},
	{
		"EE00D40C823060",
		[]int{14, 0},
	},
	{
		"8A004A801A8002F478",
		[]int{16, 315},
	},
	{
		"620080001611562C8802118E34",
		[]int{12, 315},
	},
	{
		"C0015000016115A2E0802F182340",
		[]int{23, 315},
	},
	{
		"A0016C880162017C3686B18A3D4780",
		[]int{31, 315},
	},

	{"C200B40A82", []int{14, 3}},
}

var testCasesB = []struct {
	input  string
	expect []int
}{
	{"C200B40A82", []int{0, 3}},
	{"04005AC33890", []int{0, 54}},
	{"880086C3E88112", []int{0, 7}},
	{"CE00C43D881120", []int{0, 9}},
	{"D8005AC2A8F0", []int{0, 1}},
	{"F600BC2D8F", []int{0, 0}},
	{"9C005AC2F8F0", []int{0, 0}},
	{"9C0141080250320F1802104A08", []int{0, 1}},
}

func TestSum(t *testing.T) {
	for _, c := range testCasesA {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestResult(t *testing.T) {
	for _, c := range testCasesB {
		t.Run(c.input, func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	t.Skip()
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 934, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 912901337844, B(input))
}
