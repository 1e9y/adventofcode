package cmd

import (
	"fmt"

	"github.com/1e9y/adventofcode/2015/day03"
	"github.com/1e9y/adventofcode/2015/day04"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/1e9y/adventofcode/2015/day01"
	"github.com/1e9y/adventofcode/2015/day02"
	"github.com/1e9y/adventofcode/challenge"
)

func registerEvent2015(cmd *cobra.Command) {
	year := "2015"
	event := &cobra.Command{
		Use:   year,
		Short: fmt.Sprintf("Advent of Code %d Puzzles", year),
	}
	cmd.AddCommand(event)

	registerPuzzle(event, "2015", "1", day01.A, day01.B)
	registerPuzzle(event, "2015", "2", day02.A, day02.B)
	registerPuzzle(event, "2015", "3", day03.A, day03.B)
	registerPuzzle(event, "2015", "4", day04.A, day04.B)
}

func registerPuzzle(cmd *cobra.Command, year, day string, A, B func(c *challenge.Challenge) int) {
	puzzle := &cobra.Command{
		Use:   day,
		Short: fmt.Sprintf("Puzzles for Day %s of Event %s", day, year),
	}

	puzzle.AddCommand(&cobra.Command{
		Use:   "a",
		Short: fmt.Sprintf("Day %s, Puzzle %s", day, "A"),
		Run: func(_ *cobra.Command, _ []string) {
			input := challenge.ReadChallengeForDay(year, day)
			answer := A(input)
			fmt.Printf("Answer: %d\n", answer)
		},
	})

	puzzle.AddCommand(&cobra.Command{
		Use:   "b",
		Short: fmt.Sprintf("Day %s, Puzzle %s", day, "B"),
		Run: func(_ *cobra.Command, _ []string) {
			input := challenge.ReadChallengeForDay(year, day)
			answer := B(input)
			fmt.Printf("Answer: %d\n", answer)
		},
	})

	cmd.AddCommand(puzzle)
}

var root = &cobra.Command{
	Use:     "aoc",
	Short:   "Advent of Code Solutions",
	Long:    "Golang implementation for the Advent of Code puzzles",
	Example: "go run main.go 2015 1 a -i ./2015/day01/input.txt",
	Args:    cobra.ExactArgs(1),
}

func init() {
	registerEvent2015(root)

	flags := root.PersistentFlags()
	flags.StringP("input", "i", "", "Path to the puzzle input")

	_ = viper.BindPFlags(flags)
}

func Execute() {
	if err := root.Execute(); err != nil {
		panic(err)
	}
}
