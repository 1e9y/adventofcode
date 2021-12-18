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

func parseCave(name string, index *map[string]*Cave) *Cave {
	cave := (*index)[name]
	if cave != nil {
		return cave
	}

	size := SmallCave
	if name[0] >= 'A' && name[0] <= 'Z' {
		size = BigCave
	}

	cave = &Cave{
		name:     name,
		size:     size,
		adjacent: make([]*Cave, 0),
	}
	(*index)[name] = cave
	if name == "start" {
		(*index)["start"] = cave
	}
	if name == "end" {
		(*index)["end"] = cave
	}
	return cave
}

type Path struct {
	route     []*Cave
	caves     map[string]int
	revisited bool
}

func newPath(cave *Cave) *Path {
	return &Path{
		route: []*Cave{cave},
		caves: map[string]int{cave.name: 1},
	}
}

func (path Path) last() *Cave {
	return path.route[len(path.route)-1]
}

func (path Path) append(cave *Cave) Path {
	newRoute := []*Cave{}
	newRoute = append(newRoute, path.route...)
	newRoute = append(newRoute, cave)

	newCaves := map[string]int{}
	for k, v := range path.caves {
		newCaves[k] = v
	}

	if cave.size == SmallCave {
		newCaves[cave.name]++
	}

	newPath := Path{
		route:     newRoute,
		caves:     newCaves,
		revisited: path.revisited,
	}

	if newPath.caves[cave.name] >= 2 {
		newPath.revisited = true
	}

	return newPath
}

func (path Path) canHave(cave *Cave, allowRevisit bool) bool {
	if path.caves[cave.name] >= 2 {
		return false
	}

	if path.caves[cave.name] >= 1 && (allowRevisit && path.revisited || !allowRevisit) {
		return false
	}

	return true
}

func paths(input []string, allowRevisit bool) (result int) {
	index := make(map[string]*Cave)

	for _, line := range input {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			panic(fmt.Errorf("bad input: %s", line))
		}

		a := parseCave(parts[0], &index)
		b := parseCave(parts[1], &index)
		a.adjacent = append(a.adjacent, b)
		b.adjacent = append(b.adjacent, a)
	}

	start := index["start"]
	end := index["end"]

	stack := []Path{*newPath(start)}

	for len(stack) > 0 {
		path := stack[0]
		stack = stack[1:]

		last := path.last()
		if last == end {
			result++
			continue
		}

		for _, cave := range last.adjacent {
			if cave == start {
				continue
			}

			if path.canHave(cave, allowRevisit) {
				stack = append(stack, path.append(cave))
			}
		}
	}

	return
}

func A(input *challenge.Challenge) int {
	return paths(input.LineSlice(), false)
}

func B(input *challenge.Challenge) int {
	return paths(input.LineSlice(), true)
}
