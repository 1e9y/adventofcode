package cmd

import (
	"github.com/spf13/cobra"

	"github.com/1e9y/adventofcode/2021/day01"
	"github.com/1e9y/adventofcode/2021/day02"
	"github.com/1e9y/adventofcode/2021/day03"
	"github.com/1e9y/adventofcode/2021/day05"
	"github.com/1e9y/adventofcode/2021/day06"
	"github.com/1e9y/adventofcode/2021/day07"
	"github.com/1e9y/adventofcode/2021/day08"
	"github.com/1e9y/adventofcode/2021/day09"
	"github.com/1e9y/adventofcode/2021/day10"
	"github.com/1e9y/adventofcode/2021/day11"
	"github.com/1e9y/adventofcode/2021/day12"
	"github.com/1e9y/adventofcode/2021/day13"
	"github.com/1e9y/adventofcode/2021/day14"
	"github.com/1e9y/adventofcode/2021/day15"
	"github.com/1e9y/adventofcode/2021/day16"
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
		registerPuzzle(event, year, "8", day08.A, day08.B)
		registerPuzzle(event, year, "9", day09.A, day09.B)
		registerPuzzle(event, year, "10", day10.A, day10.B)
		registerPuzzle(event, year, "11", day11.A, day11.B)
		registerPuzzle(event, year, "12", day12.A, day12.B)
		registerPuzzle(event, year, "13", day13.A, day13.B)
		registerPuzzle(event, year, "14", day14.A, day14.B)
		registerPuzzle(event, year, "15", day15.A, day15.B)
		registerPuzzle(event, year, "16", day16.A, day16.B)
	})
}
