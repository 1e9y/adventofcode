package day02

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type CubeType string

const (
	Red   CubeType = "red"
	Green CubeType = "green"
	Blue  CubeType = "blue"
)

var gameIDRe = regexp.MustCompile(`^Game (\d+)`)

type Game map[CubeType]int
type Configuration map[CubeType]int

var DefaultConfiguration Configuration = map[CubeType]int{
	Red:   12,
	Green: 13,
	Blue:  14,
}

func parse(input string) (id int, games []Game) {
	parts := gameIDRe.FindStringSubmatch(input)
	if len(parts) < 2 {
		panic(fmt.Sprintf("bad input: %s", input))
	}
	id = util.MustAtoi(parts[1])

	parts = strings.Split(input, ": ")
	if len(parts) < 2 {
		panic(fmt.Sprintf("bad input: %s", input))
	}

	return id, parseGames(parts[1])
}

func parseGames(input string) (round []Game) {
	parts := strings.Split(input, "; ")
	for _, game := range parts {
		round = append(round, parseGame(game))
	}
	return
}

func parseGame(input string) (game Game) {
	game = make(map[CubeType]int)
	parts := strings.Split(input, ", ")
	for _, part := range parts {
		pick := strings.Split(part, " ")
		if len(pick) < 2 {
			panic(fmt.Sprintf("bad input: %s", input))
		}
		count := util.MustAtoi(pick[0])
		color := pick[1]
		game[CubeType(color)] = count
	}
	return
}

func valid(games []Game) bool {
	for _, game := range games {
		for color, count := range game {
			if count > DefaultConfiguration[color] {
				return false
			}
		}
	}
	return true
}

func power(games []Game) int {
	minConfiguration := make(Configuration)
	for _, game := range games {
		for color, count := range game {
			minConfiguration[color] = util.MaxInt(count, minConfiguration[color])
		}
	}
	result := 1
	for _, v := range minConfiguration {
		result *= v
	}
	return result
}

func A(input *challenge.Challenge) int {
	result := 0
	for line := range input.Lines() {
		id, games := parse(line)
		if valid(games) {
			result += id
		}
	}
	return result
}

func B(input *challenge.Challenge) int {
	result := 0
	for line := range input.Lines() {
		_, games := parse(line)
		result += power(games)
	}
	return result
}
