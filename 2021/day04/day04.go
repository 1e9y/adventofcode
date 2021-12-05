package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

type Board struct {
	grid       [5][5]int
	markedRows map[int]int
	markedCols map[int]int
	completed  bool
	done       bool
}

func (b *Board) Completed() bool {
	return b.completed
}

func (b *Board) Done() bool {
	return b.done
}

func (c *Board) Mark(n int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if c.grid[i][j] == n {
				c.markedRows[i]++
				c.markedCols[j]++
				if c.markedRows[i] == 5 || c.markedCols[j] == 5 {
					c.completed = true
					// fmt.Println("mark", n, c)
				}
			}
		}
	}
}

func (c *Board) Score(won int, nums map[int]bool) int {
	sum := 0
	// fmt.Println(nums)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if _, ok := nums[c.grid[i][j]]; !ok {
				// fmt.Println(c.grid[i][j])
				sum += c.grid[i][j]
			}
		}
	}

	return sum * won
}

func newBoardFromInput(input <-chan string) (*Board, bool) {
	b := Board{
		grid:       [5][5]int{},
		markedRows: make(map[int]int),
		markedCols: make(map[int]int),
		completed:  false,
		done:       false,
	}

	for i := 0; i < 5; i++ {
		line, ok := <-input
		if !ok {
			return nil, false
		}
		parts := strings.Split(line, " ")
		row := [5]int{}
		p := 0
		for _, s := range parts {
			n, err := strconv.Atoi(s)
			if err == nil {
				row[p] = n
				p++
			}
		}
		b.grid[i] = row
	}

	return &b, true
}

func bingo(input <-chan string) (int, int) {
	calledNumber := make(map[int]bool)
	rawNumbers := <-input
	var numbers []int
	for _, p := range strings.Split(rawNumbers, ",") {
		numbers = append(numbers, util.MustAtoi(p))
	}
	// fmt.Println(rawNumbers)
	// fmt.Println(numbers)
	<-input

	var boards []*Board

	for {
		b, ok := newBoardFromInput(input)
		if !ok {
			break
		}
		boards = append(boards, b)
		<-input
		// fmt.Println(b)
	}

	var first, last int
	// game
	for _, n := range numbers {
		calledNumber[n] = true
		for _, b := range boards {
			if b.Completed() {
				continue
			}
			b.Mark(n)
			if b.Completed() {
				// fmt.Println("WON")
				// fmt.Println(b)
				if first == 0 {
					first = b.Score(n, calledNumber)
					fmt.Println("GOT THE FIST WINNER")
				}
				// boards[i] = boards[len(boards)-1]
				// boards = boards[:len(boards)-1]
				last = b.Score(n, calledNumber)
				fmt.Println("GOT THE NEXT WINNER", last)
			}
		}
	}

	fmt.Println("BAD....", first, last)

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
