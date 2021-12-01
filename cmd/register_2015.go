package cmd

import (
	"github.com/spf13/cobra"

	"github.com/1e9y/adventofcode/2015/day01"
	"github.com/1e9y/adventofcode/2015/day02"
	"github.com/1e9y/adventofcode/2015/day03"
	"github.com/1e9y/adventofcode/2015/day04"
)

func registerEvent2015(cmd *cobra.Command) {
	year := "2015"
	registerEvent(cmd, year, func(event *cobra.Command) {
		registerPuzzle(event, year, "1", day01.A, day01.B)
		registerPuzzle(event, year, "2", day02.A, day02.B)
		registerPuzzle(event, year, "4", day04.A, day04.B)
		registerPuzzle(event, year, "3", day03.A, day03.B)
	})
}
