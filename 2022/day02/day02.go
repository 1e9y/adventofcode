package day02

import (
	"fmt"
	"github.com/1e9y/adventofcode/challenge"
	"strings"
)

type Shape rune

const (
	OpponentRock     Shape = 'A'
	OpponentPaper          = 'B'
	OpponentScissors       = 'C'

	PlayerRock     Shape = 'X'
	PlayerPaper          = 'Y'
	PlayerScissors       = 'Z'

	PlayerLoss Shape = 'X'
	PlayerDraw       = 'Y'
	PlayerWin        = 'Z'

	ScoreWin  = 6
	ScoreDraw = 3
	ScoreLoss = 0
)

var ShapeScores = map[Shape]int{
	PlayerRock:     1,
	PlayerPaper:    2,
	PlayerScissors: 3,
}

var OutcomesMap = map[Shape]map[Shape]int{
	OpponentRock: {
		PlayerRock:     ScoreDraw,
		PlayerPaper:    ScoreWin,
		PlayerScissors: ScoreLoss,
	},
	OpponentPaper: {
		PlayerRock:     ScoreLoss,
		PlayerPaper:    ScoreDraw,
		PlayerScissors: ScoreWin,
	},
	OpponentScissors: {
		PlayerRock:     ScoreWin,
		PlayerPaper:    ScoreLoss,
		PlayerScissors: ScoreDraw,
	},
}

var PlayerChoiceMap = map[Shape]map[Shape]Shape{
	PlayerLoss: {
		OpponentRock:     PlayerScissors,
		OpponentPaper:    PlayerRock,
		OpponentScissors: PlayerPaper,
	},
	PlayerDraw: {
		OpponentRock:     PlayerRock,
		OpponentPaper:    PlayerPaper,
		OpponentScissors: PlayerScissors,
	},
	PlayerWin: {
		OpponentRock:     PlayerPaper,
		OpponentPaper:    PlayerScissors,
		OpponentScissors: PlayerRock,
	},
}

func parseRound(input string) []Shape {
	parts := strings.Split(input, " ")

	if parts[0][0] < 'A' || parts[0][0] > 'C' ||
		parts[1][0] < 'X' || parts[1][0] > 'Z' {
		panic(fmt.Errorf("bad input: %v", input))
	}

	return []Shape{
		Shape(parts[0][0]),
		Shape(parts[1][0]),
	}
}

func parseRound2(input string) []Shape {
	parts := strings.Split(input, " ")

	if parts[0][0] < 'A' || parts[0][0] > 'C' ||
		parts[1][0] < 'X' || parts[1][0] > 'Z' {
		panic(fmt.Errorf("bad input: %v", input))
	}

	opponent := Shape(parts[0][0])
	player := PlayerChoiceMap[Shape(parts[1][0])][opponent]

	return []Shape{
		opponent,
		player,
	}
}

func score(round []Shape) int {
	opponent := round[0]
	player := round[1]

	return OutcomesMap[opponent][player] + ShapeScores[round[1]]
}

func tournament(input <-chan string, parse func(string) []Shape) int {
	total := 0
	for line := range input {
		total += score(parse(line))
	}
	return total
}

func A(input *challenge.Challenge) int {
	return tournament(input.Lines(), parseRound)
}

func B(input *challenge.Challenge) int {
	return tournament(input.Lines(), parseRound2)
}
