package challenge

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/1e9y/adventofcode/util"

	"github.com/spf13/viper"
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

func ReadChallengeForDay(year, day string) *Challenge {
	path := viper.GetString("input")
	if path == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		path = filepath.Join(wd, year, fmt.Sprintf("day%02d", util.MustAtoi(day)), "input.txt")
	}

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return newChallengeFromReader(file, file)
}

func ReadChallengeFromFile() *Challenge {
	path := viper.GetString("input")
	if path == "" {
		_, f, _, ok := runtime.Caller(1)
		if !ok {
			panic("failed to determine input path, provide it with -i explicitly")
		}

		path = filepath.Join(filepath.Dir(f), "input.txt")
		println(f)
	}

	file, err := os.Open(path)
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
