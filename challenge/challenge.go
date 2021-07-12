package challenge

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Challenge struct {
	scanner *bufio.Scanner
	lines   chan string
}

func newChallengeFromReader(r io.Reader, c io.Closer) *Challenge {
	challenge := &Challenge{
		scanner: bufio.NewScanner(r),
		lines:   make(chan string),
	}

	go func() {
		defer func() {
			if c != nil {
				_ = c.Close()
			}
		}()

		for challenge.scanner.Scan() {
			challenge.lines <- challenge.scanner.Text()
		}

		close(challenge.lines)
	}()

	return challenge
}

//func ReadChallengeForDay(day int) *Challenge {}

func ReadChallengeFromFile(path string) *Challenge {
	p := path
	if path == "" {
		_, f, _, ok := runtime.Caller(1)
		if !ok {
			panic("failed to determine input path")
		}

		p = filepath.Join(filepath.Dir(f), "input.txt")
		println(f)
	}

	file, err := os.Open(p)
	println(p)
	if err != nil {
		panic(err)
	}

	return newChallengeFromReader(file, file)
}

func ReadChallengeFromLiteral(input string) *Challenge {
	return newChallengeFromReader(strings.NewReader(input), nil)
}

func (c *Challenge) Lines() <-chan string {
	return c.lines
}

func (c *Challenge) LineSlice() (result []string) {
	for line := range c.lines {
		result = append(result, line)
	}
	return
}
