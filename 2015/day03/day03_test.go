package day03

import (
	"testing"

	"github.com/1e9y/adventofcode/challenge"
)

func TestA(t *testing.T) {
	input := challenge.ReadChallengeFromLiteral("what is it")
	got := a(input)
	if got != 124 {
		t.Fatalf("got: %v, want: 124", got)
	}
}

func TestInput(t *testing.T) {
	input := challenge.ReadChallengeFromFile("./day03.input")
	got := a(input)
	if got != 124 {
		t.Fatalf("got: %v, want: 124", got)
	}
}

func Register(s *S.Re)
