package day12

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
		`start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
		[]int{10, 36},
	},
	{
		`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`,
		[]int{19, 103},
	},
	{
		`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`,
		[]int{226, 3509},
	},
}

func TestPaths(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:12], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[0], A(input))
		})
	}
}

func TestTwice(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.input[:12], func(t *testing.T) {
			input := challenge.ReadChallengeFromLiteral(c.input)
			assert.Equal(t, c.expect[1], B(input))
		})
	}
}

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 4970, A(input))
}

func TestB(t *testing.T) {
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 137948, B(input))
}
