package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/1e9y/adventofcode/challenge"
)

type Solution struct {
	A func(c *challenge.Challenge) int
	B func(c *challenge.Challenge) int
}

func registerEvent(cmd *cobra.Command, year string, puzzles func(event *cobra.Command)) {
	event := &cobra.Command{
		Use:   year,
		Short: fmt.Sprintf("Advent of Code %s puzzles", year),
	}
	cmd.AddCommand(event)
	puzzles(event)
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
	registerEvent2021(root)

	flags := root.PersistentFlags()
	flags.StringP("input", "i", "", "Path to the puzzle input")

	_ = viper.BindPFlags(flags)
}

func Execute() {
	if err := root.Execute(); err != nil {
		panic(err)
	}
}
