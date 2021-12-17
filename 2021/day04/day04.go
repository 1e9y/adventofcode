package day04

import (
	"strconv"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type Bingo struct {
	board      [5][5]int
	markedRows map[int]int
	markedCols map[int]int
	completed  bool
}

func newBingoFromInput(input <-chan string) (*Bingo, bool) {
	bingo := Bingo{
		board:      [5][5]int{},
		markedRows: make(map[int]int),
		markedCols: make(map[int]int),
		completed:  false,
	}

	for i := 0; i < 5; i++ {
		line := <-input
		if line == "" {
			return nil, false
		}
		parts := strings.Split(line, " ")
		row := [5]int{}
		j := 0
		for _, s := range parts {
			n, err := strconv.Atoi(s)
			if err == nil {
				row[j] = n
				j++
			}
		}
		bingo.board[i] = row
	}

	return &bingo, true
}

func (b *Bingo) Completed() bool {
	return b.completed
}

func (c *Bingo) Mark(n int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if c.board[i][j] == n {
				c.markedRows[i]++
				c.markedCols[j]++
				if c.markedRows[i] == 5 || c.markedCols[j] == 5 {
					c.completed = true
				}
			}
		}
	}
}

func (c *Bingo) Score(won int, nums map[int]bool) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if _, ok := nums[c.board[i][j]]; !ok {
				sum += c.board[i][j]
			}
		}
	}

	return sum * won
}

func bingo(input <-chan string) (int, int) {
	calledNumber := make(map[int]bool)
	rawNumbers := <-input
	var numbers []int
	for _, p := range strings.Split(rawNumbers, ",") {
		numbers = append(numbers, util.MustAtoi(p))
	}
	<-input

	var bingos []*Bingo

	for {
		bingo, ok := newBingoFromInput(input)
		if !ok {
			break
		}
		bingos = append(bingos, bingo)
		<-input
	}

	var first, last int
	for _, n := range numbers {
		calledNumber[n] = true
		for _, bingo := range bingos {
			if bingo.Completed() {
				continue
			}
			bingo.Mark(n)
			if bingo.Completed() {
				if first == 0 {
					first = bingo.Score(n, calledNumber)
				}
				last = bingo.Score(n, calledNumber)
			}
		}
	}

	return first, last
}

func A(input *challenge.Challenge) int {
	answer, _ := bingo(input.Lines())
	return answer
}

func B(input *challenge.Challenge) int {
	_, answer := bingo(input.Lines())
	return answer
}
