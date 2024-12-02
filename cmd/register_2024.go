package cmd

import (
	"github.com/spf13/cobra"

	"github.com/1e9y/adventofcode/2024/day01"
)

func registerEvent2024(cmd *cobra.Command) {
	year := "2024"
	registerEvent(cmd, year, func(event *cobra.Command) {
		registerPuzzle(event, year, "1", day01.A, day01.B)
	})
}
