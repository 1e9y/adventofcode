package cmd

import (
	"github.com/spf13/cobra"

	"github.com/1e9y/adventofcode/2015/day02"
	"github.com/1e9y/adventofcode/2021/day01"
)

func registerEvent2021(cmd *cobra.Command) {
	year := "2021"
	registerEvent(cmd, year, func(event *cobra.Command) {
		registerPuzzle(event, year, "1", day01.A, day01.B)
		registerPuzzle(event, year, "2", day02.A, day02.B)
	})
}
