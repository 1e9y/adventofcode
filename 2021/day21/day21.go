package day21

import (
	"fmt"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type Dice struct {
	sides int
	value int
	rolls int
}

func (dice *Dice) roll() int {
	dice.rolls++
	dice.value = dice.value%dice.sides + 1
	return dice.value
}

type DeterministicGame struct {
	turn      int
	positions []int
	scores    []int
	dice      Dice
}

func (game *DeterministicGame) play() bool {
	a := game.dice.roll()
	b := game.dice.roll()
	c := game.dice.roll()

	turn := game.turn
	game.positions[turn] = (game.positions[turn]+(a+b+c)-1)%10 + 1
	game.scores[turn] += game.positions[turn]

	if game.scores[turn] >= 1000 {
		return false
	}

	game.turn = game.turn ^ 1
	return true
}

func deterministic(pos1, pos2 int) int {
	game := DeterministicGame{
		turn:      0,
		positions: []int{pos1, pos2},
		scores:    make([]int, 2),
		dice:      Dice{sides: 100},
	}

	for game.play() {
	}

	return util.MinInt(game.scores[0], game.scores[1]) * game.dice.rolls
}

type DiracGame struct {
	turn      int
	positions []int
	scores    []int64
}

func (game DiracGame) String() string {
	return fmt.Sprintf("%d:%d:%d:%d:%d", game.positions[0], game.scores[0], game.positions[1], game.scores[1], game.turn)
}

var diceDistribution = map[int]int64{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func dirac(pos1, pos2 int) int64 {
	cache := make(map[string]int64)

	game := DiracGame{
		turn:      0,
		positions: []int{pos1, pos2},
		scores:    []int64{0, 0},
	}

	var play func(game DiracGame) int64
	play = func(game DiracGame) int64 {
		if _, ok := cache[game.String()]; ok {
			return cache[game.String()]
		}

		var sum int64 = 0
		for dice, universes := range diceDistribution {
			newGame := DiracGame{
				turn:      game.turn,
				positions: []int{game.positions[0], game.positions[1]},
				scores:    []int64{game.scores[0], game.scores[1]},
			}

			newGame.positions[newGame.turn] = (newGame.positions[newGame.turn]+dice-1)%10 + 1
			newGame.scores[newGame.turn] += int64(newGame.positions[newGame.turn])

			if newGame.scores[newGame.turn] >= 21 {
				if newGame.turn == 0 {
					sum += universes
				}
			} else {
				newGame.turn = newGame.turn ^ 1
				sum += universes * play(newGame)
			}
		}

		cache[game.String()] = sum
		return sum
	}

	return play(game)
}

func parseInput(input []string) (int, int) {
	return int(input[0][len(input[0])-1] - '0'), int(input[1][len(input[1])-1] - '0')
}

func A(input *challenge.Challenge) int {
	return deterministic(parseInput(input.LineSlice()))
}

func B(input *challenge.Challenge) int {
	return int(dirac(parseInput(input.LineSlice())))
}
