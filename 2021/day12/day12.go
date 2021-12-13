package day12

import (
	"fmt"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
)

type CaveSize byte

const (
	SmallCave CaveSize = iota
	BigCave
)

type Cave struct {
	name     string
	size     CaveSize
	adjacent []*Cave
}

var start, end *Cave
var index = make(map[string]*Cave)

func caveSize(name string) CaveSize {
	if name[0] >= 'A' && name[0] <= 'Z' {
		return BigCave
	}
	return SmallCave
}

func parseCave(name string) *Cave {
	cave := index[name]
	if cave != nil {
		return cave
	}
	cave = &Cave{
		name:     name,
		size:     caveSize(name),
		adjacent: make([]*Cave, 0),
	}
	index[name] = cave
	if name == "start" {
		start = cave
	}
	if name == "end" {
		end = cave
	}
	return cave
}

func printPath(path []*Cave) {
	res := make([]string, 0)
	for _, p := range path {
		res = append(res, p.name)
	}
	fmt.Println(strings.Join(res, ","))
}

func pathHas(path []*Cave, cave *Cave) bool {
	for _, c := range path {
		if c.name == cave.name {
			return true
		}
	}

	return false
}

func canHave(path []*Cave, cave *Cave) bool {
	if cave.size == BigCave {
		return true
	}
	index := make(map[string]int)
	hasTwo := 0
	for _, c := range path {
		index[c.name]++
		if c.size == BigCave {
			continue
		}
		if index[c.name] >= 2 {
			// if hasTwo {
			// return false
			// }
			hasTwo += 1
			if hasTwo == 2 {
				return false
			}
			// return false
		}
	}

	// if hasTwo || index[cave.name] >= 2 {

	return true
}

func paths(input []string) (result int) {
	for _, line := range input {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			panic("bad input")
		}

		a := parseCave(parts[0])
		b := parseCave(parts[1])
		a.adjacent = append(a.adjacent, b)
		b.adjacent = append(b.adjacent, a)
	}

	// fmt.Println(index)

	completedPaths := make([][]*Cave, 0)
	stack := make([][]*Cave, 0)
	stack = append(stack, []*Cave{start})

	for len(stack) > 0 {
		// println("!", len(stack))
		// for _, p := range stack {
		// 	printPath(p)
		// }
		path := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// path := stack[0]
		// stack = stack[1:]

		// printPath(path)
		last := path[len(path)-1]
		if last == end {
			// println("> done!")
			completedPaths = append(completedPaths, path)
			continue
		}

		// print("+ ")
		// printPath(last.adjacent)

		for _, a := range last.adjacent {
			if a == start {
				continue
			}
			// if a.size == BigCave || !pathHas(path, a) {

			// fmt.Printf("+ checking if path can have %v:\n", a.name)
			// printPath(path)
			if canHave(path, a) {
				// println("+ YES")
				// TODO: What was wrong with simple nextPath?
				nextPath := make([]*Cave, 0)
				nextPath = append(nextPath, path...)
				nextPath = append(nextPath, a)
				stack = append(stack, nextPath)
				// fmt.Printf("= adding %v to %v\n", a.name, nextPath)
				// fmt.Printf("= new stack\n")
				// for _, p := range stack {
				// 	printPath(p)
				// }
				// fmt.Printf("=\n")
			}
		}
	}

	// for _, p := range completedPaths {
	// 	printPath(p)
	// }
	return len(completedPaths)
}

func A(input *challenge.Challenge) int {
	return paths(input.LineSlice())
}

func B(input *challenge.Challenge) int {
	return paths(input.LineSlice())
}
