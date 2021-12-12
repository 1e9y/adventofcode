package day10

import (
	"sort"

	"github.com/1e9y/adventofcode/challenge"
)

var BracketPair = map[rune]rune{
	'(': ')',
	'[': ']',
	'<': '>',
	'{': '}',
}

var BracketSyntaxScores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var BracketCompletenessScores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func syntaxScore(input []rune) int {
	return BracketSyntaxScores[input[0]]
}

func completenessScore(input []rune) (result int) {
	for i := len(input) - 1; i >= 0; i-- {
		result = result*5 + BracketCompletenessScores[input[i]]
	}
	return
}

func validate(input string) (bool, []rune) {
	var expected []rune
	for _, r := range input {
		if p, ok := BracketPair[r]; ok {
			expected = append(expected, p)
		} else {
			if r == expected[len(expected)-1] {
				expected = expected[:len(expected)-1]
			} else {
				return false, []rune{r}
			}
		}
	}

	return true, expected
}

func score(input []string) (syntax, compleness int) {
	scores := make([]int, 0)
	for _, line := range input {
		if ok, bad := validate(line); ok {
			scores = append(scores, completenessScore(bad))
		} else {
			syntax += syntaxScore(bad)
		}
	}
	sort.Ints(scores)
	compleness = scores[len(scores)/2]
	return
}

func A(input *challenge.Challenge) int {
	answer, _ := score(input.LineSlice())
	return answer
}

func B(input *challenge.Challenge) int {
	_, answer := score(input.LineSlice())
	return answer
}
