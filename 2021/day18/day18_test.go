package day18

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/stretchr/testify/assert"
)

// var testCases = []struct {
// 	input  string
// 	expect []int
// }{
// 	{
// 		"target area: x=20..30, y=-10..-5",
// 		[]int{45, 112},
// 	},
// }

var testCasesStrings = []string{

	"[[[99,[3,8]],5],1]",
	"[1,2]",
	"[[1,2],3]",
	"[9,[8,7]]",
	"[[1,9],[8,5]]",
	"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]",
	"[[[9,[3,8]],[[0,9],6]],[[[3,7],[4,9]],3]]",
	"[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]",
}

func TestString(t *testing.T) {
	// t.Skip()
	for _, input := range testCasesStrings {
		t.Run(input, func(t *testing.T) {
			number := newNumberFromInput(input)
			assert.Equal(t, input, number.String())
		})
	}
}

var testCasesAddition = []struct {
	input  []string
	expect string
}{
	// {
	// 	[]string{"[1,2]", "[[3,4],5]"},
	// 	"[[1,2],[[3,4],5]]",
	// },
	// {
	// 	[]string{"[[1,2],3]", "[9,[8,7]]"},
	// 	"[[[1,2],3],[9,[8,7]]]",
	// },
	// {
	// 	[]string{"[1,9]", "[8,5]"},
	// 	"[[1,9],[8,5]]",
	// },

	/*
		   [0
		   	[1
		   		[
		   			[
		   				4,
		   				0
		   			],
		   			[
		   				5,
		   				0
		   			]
		   		],
		   		[
		   			[
		   				[
		   					4,
							   5],
		   				[2,6]
		   			],
		   			[9,5]
		   		]
		   	]
		   	[
		   		7,
		   		[
		   			[
		   				[3,7],
		   				[4,3]
		   			],
		   			[
		   				[6,3],
		   				[8,8]
		   			]
		   		]
		   	]
		   ]
	*/

	{
		[]string{
			"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
		},
		"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
	},
}

func TestAddition(t *testing.T) {
	// t.Skip()
	for _, c := range testCasesAddition {
		t.Run(c.input[0], func(t *testing.T) {
			assert.Equal(t,
				c.expect,
				add(newNumberFromInput(c.input[0]), newNumberFromInput(c.input[1])).String(),
			)
		})
	}
}

var testCasesSplit = []struct {
	input  string
	expect string
}{
	{
		"[[[[0,7],4],[15,[0,13]]],[1,1]]",
		"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
	},
	{
		"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]",
		"[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]",
	},
}

func TestSplit(t *testing.T) {
	// t.Skip()
	for _, c := range testCasesSplit {
		t.Run(c.input, func(t *testing.T) {
			number := newNumberFromInput(c.input)
			_ = number.split()
			assert.Equal(t, c.expect, number.String())
		})
	}
}

var testCasesExplode = []struct {
	input  string
	expect string
}{
	{
		"[[[[[9,8],1],2],3],4]",
		"[[[[0,9],2],3],4]",
	},
	{
		"[7,[6,[5,[4,[3,2]]]]]",
		"[7,[6,[5,[7,0]]]]",
	},
	{
		"[[6,[5,[4,[3,2]]]],1]",
		"[[6,[5,[7,0]]],3]",
	},
	{
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
	},
	{
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	},
	{
		"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
		"[[[[0,7],4],[7,[[8,4],9]]],[1,1]]",
	},
}

func TestExplode(t *testing.T) {
	// t.Skip()
	for _, c := range testCasesExplode {
		t.Run(c.input, func(t *testing.T) {
			number := newNumberFromInput(c.input)
			_ = number.explode()
			assert.Equal(t, c.expect, number.String())
		})
	}
}

var testCasesReduce = []struct {
	input  []string
	expect string
}{
	{
		[]string{
			"[[[[4,3],4],4],[7,[[8,4],9]]]",
			"[1,1]",
		},
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
	},
	{
		[]string{
			"[1,1]",
			"[2,2]",
			"[3,3]",
			"[4,4]",
		},
		"[[[[1,1],[2,2]],[3,3]],[4,4]]",
	},
	{
		[]string{
			"[1,1]",
			"[2,2]",
			"[3,3]",
			"[4,4]",
			"[5,5]",
		},
		"[[[[3,0],[5,3]],[4,4]],[5,5]]",
	},
	{
		[]string{
			"[1,1]",
			"[2,2]",
			"[3,3]",
			"[4,4]",
			"[5,5]",
			"[6,6]",
		},
		"[[[[5,0],[7,4]],[5,5]],[6,6]]",
	},
	{
		[]string{
			"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
			"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
			"[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]",
			"[7,[5,[[3,8],[1,4]]]]",
			"[[2,[2,2]],[8,[8,1]]]",
			"[2,9]",
			"[1,[[[9,3],9],[[9,0],[0,7]]]]",
			"[[[5,[7,4]],7],1]",
			"[[[[4,2],2],6],[8,7]]",
		},
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
	},
}

func TestReduce(t *testing.T) {
	// t.Skip()
	for _, c := range testCasesReduce {
		t.Run(c.input[0], func(t *testing.T) {
			a := newNumberFromInput(c.input[0])
			for _, s := range c.input[1:] {
				b := newNumberFromInput(s)
				a = add(a, b)
				// a.reduce()
			}
			// a.reduce()
			assert.Equal(t, c.expect, a.String())
		})
	}
}

var testCasesMagnitude = []struct {
	input  string
	expect int
}{
	{
		"[9,1]", 29,
	},

	{"[[1,2],[[3,4],5]]", 143},
	{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
	{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
	{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
	{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
	{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
}

func TestMagnitude(t *testing.T) {
	// t.Skip()
	for _, c := range testCasesMagnitude {
		t.Run(c.input, func(t *testing.T) {
			number := newNumberFromInput(c.input)
			assert.Equal(t, c.expect, number.magnitude())
		})
	}
}

var testCasesHomework = []struct {
	input  string
	expect []int
}{
	{
		`[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`,
		[]int{4140, 3993},
	},
}

func TestHomework(t *testing.T) {
	// t.Skip()
	for _, c := range testCasesHomework {
		input := challenge.ReadChallengeFromLiteral(c.input)
		assert.Equal(t, c.expect[0], A(input))
	}
}
func TestSecondHomework(t *testing.T) {
	// t.Skip()
	for _, c := range testCasesHomework {
		input := challenge.ReadChallengeFromLiteral(c.input)
		assert.Equal(t, c.expect[1], B(input))
	}
}

// func TestString(t *testing.T) {
// 	for _, c := range testCasesStrings {
// 		t.Run(c.input, func(t *testing.T) {
// 			input := challenge.ReadChallengeFromLiteral(c.input)
// 			assert.Equal(t, c.expect[0], A(input))
// 		})
// 	}
// }

func TestA(t *testing.T) {
	// t.Skip()
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 4132, A(input))
}

func TestB(t *testing.T) {
	// t.Skip()
	input := challenge.ReadChallengeFromFile()
	assert.Equal(t, 4685, B(input))
}
