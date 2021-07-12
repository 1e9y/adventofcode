package day01

import (
	"fmt"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/spf13/cobra"
)

func stairs(input string) (int, int) {
	f := 0
	b := 0
	for i, c := range input {
		if c == '(' {
			f++
		} else {
			f--
		}
		if b == 0 && f == -1 {
			b = i + 1
		}
	}
	return f, b
}

func a(input *challenge.Challenge) int {
	var answer int
	answer, _ = stairs(input.LineSlice()[0])
	return answer
}

func b(input *challenge.Challenge) int {
	var answer int
	_, answer = stairs(input.LineSlice()[0])
	return answer
}

func Register(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "1",
		Short: "Problems for Day 1",
	}

	day.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Day 1, Problem A",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", a(challenge.ReadChallengeFromFile("")))
		},
	})

	day.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Day 1, Problem B",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("Answer: %d\n", b(challenge.ReadChallengeFromFile("")))
		},
	})

	root.AddCommand(day)
}
