package cmd

import (
	"github.com/spf13/cobra"

	"github.com/1e9y/adventofcode/2023/day01"
)

func registerEvent2023(cmd *cobra.Command) {
	year := "2023"
	registerEvent(cmd, year, func(event *cobra.Command) {
		registerPuzzle(event, year, "1", day01.A, day01.B)
		// registerPuzzle(event, year, "2", day02.A, day02.B)
		// registerPuzzle(event, year, "3", day03.A, day03.B)
		// registerPuzzle(event, year, "4", day04.A, day04.B)
		// registerPuzzle(event, year, "5", day05.A, day05.B)
		// registerPuzzle(event, year, "6", day06.A, day06.B)
	})
}
