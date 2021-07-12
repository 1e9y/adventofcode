package cmd

import (
	"github.com/1e9y/adventofcode/2015/day01"

	"github.com/spf13/cobra"
)

func registerDays(root *cobra.Command) {
	day01.Register(root)
}

func NewRootCommand() *cobra.Command {
	//var (
	//	start time.Time
	//)

	result := &cobra.Command{
		Use:     "aoc",
		Short:   "Advent of Code Solutions",
		Long:    "Goland implementation for the Advent of Code problems",
		Example: "go run main.go 2015 1 a -i ./2015/day01/input.txt",
		Args:    cobra.ExactArgs(1),
	}

	registerDays(result)

	return result
}
