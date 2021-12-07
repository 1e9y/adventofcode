package cmd

import (
	"github.com/spf13/cobra"

	"github.com/1e9y/adventofcode/2021/day01"
	"github.com/1e9y/adventofcode/2021/day02"
	"github.com/1e9y/adventofcode/2021/day03"
	"github.com/1e9y/adventofcode/2021/day05"
	"github.com/1e9y/adventofcode/2021/day06"
	"github.com/1e9y/adventofcode/2021/day07"
)

func registerEvent2021(cmd *cobra.Command) {
	year := "2021"
	registerEvent(cmd, year, func(event *cobra.Command) {
		registerPuzzle(event, year, "1", day01.A, day01.B)
		registerPuzzle(event, year, "2", day02.A, day02.B)
		registerPuzzle(event, year, "3", day03.A, day03.B)
		registerPuzzle(event, year, "5", day05.A, day05.B)
		registerPuzzle(event, year, "6", day06.A, day06.B)
		registerPuzzle(event, year, "7", day07.A, day07.B)
	})
}
