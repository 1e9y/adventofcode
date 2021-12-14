package day14

import (
	"fmt"
	"math"
	"regexp"

	"github.com/1e9y/adventofcode/challenge"
	"github.com/1e9y/adventofcode/util"
)

var templateRe = regexp.MustCompile(`^[A-Z]+$`)
var ruleRe = regexp.MustCompile(`^([A-Z]{2}) -> ([A-Z])$`)

type Rule struct {
	pair      string
	insertion byte
}

type Formula struct {
	template string
	rules    []Rule
	elements map[byte]int
	pairs    map[string]int
}

func newFormulaFromInput(input <-chan string) *Formula {
	formula := Formula{
		elements: make(map[byte]int),
		pairs:    make(map[string]int),
	}
	for line := range input {
		if templateRe.MatchString(line) {
			formula.template = line
		}
		if ruleRe.MatchString(line) {
			matches := ruleRe.FindStringSubmatch(line)
			if len(matches) != 3 {
				panic(fmt.Errorf("bad input: %s", line))
			}
			formula.rules = append(formula.rules, Rule{matches[1], matches[2][0]})
		}
	}

	template := formula.template
	formula.elements[template[0]]++
	var key string
	for i := 1; i < len(template); i++ {
		formula.elements[template[i]]++
		key = fmt.Sprintf("%c%c", template[i-1], template[i])
		formula.pairs[key]++
	}

	return &formula
}

func (formula *Formula) minmax() (int, int) {
	min := math.MaxInt64
	max := 0
	for _, count := range formula.elements {
		max = util.MaxInt(max, count)
		min = util.MinInt(min, count)
	}
	return min, max
}

type processingPair struct {
	pair  string
	count int
}

func (formula *Formula) process() {
	destroyedPairs := []processingPair{}
	createdPairs := []processingPair{}
	for _, rule := range formula.rules {
		left := fmt.Sprintf("%c%c", rule.pair[0], rule.insertion)
		right := fmt.Sprintf("%c%c", rule.insertion, rule.pair[1])
		if formula.pairs[rule.pair] > 0 {
			destroy := processingPair{rule.pair, formula.pairs[rule.pair]}
			destroyedPairs = append(destroyedPairs, destroy)

			createLeft := processingPair{left, formula.pairs[rule.pair]}
			createRight := processingPair{right, formula.pairs[rule.pair]}
			createdPairs = append(createdPairs, createLeft, createRight)

			formula.elements[rule.insertion] += formula.pairs[rule.pair]
		}
	}
	for _, pair := range destroyedPairs {
		formula.pairs[pair.pair] -= pair.count
	}
	for _, pair := range createdPairs {
		formula.pairs[pair.pair] += pair.count
	}
}

func polymerize(formula *Formula, steps int) int {
	for s := 0; s < steps; s++ {
		formula.process()
	}

	min, max := formula.minmax()
	return max - min
}

func A(input *challenge.Challenge) int {
	return polymerize(newFormulaFromInput(input.Lines()), 10)
}

func B(input *challenge.Challenge) int {
	return polymerize(newFormulaFromInput(input.Lines()), 40)
}
