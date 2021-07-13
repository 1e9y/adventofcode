package day02

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
)

func min(a ...int) int {
	m := a[0]
	for _, n := range a {
		if n < m {
			m = n
		}
	}
	return m
}

func A(input *challenge.Challenge) int {
	return 0
}

func B(input *challenge.Challenge) int {
	return 0
}

func wrapping(l, w, h int) int {
	return 2*(l*w+w*h+h*l) + min(l*w, w*h, h*l)
}

func ribbon(l, w, h int) int {
	return l*w*h + min(2*(l+w), 2*(w+h), 2*(h+l))
}

func chall1(input *os.File) (int, int) {
	scanner := bufio.NewScanner(input)
	sum := 0
	rib := 0
	for scanner.Scan() {
		t := scanner.Text()
		d := strings.Split(t, "x")
		l, _ := strconv.Atoi(d[0])
		w, _ := strconv.Atoi(d[1])
		h, _ := strconv.Atoi(d[2])
		sum += wrapping(l, w, h)
		rib += ribbon(l, w, h)
	}
	return sum, rib
}
