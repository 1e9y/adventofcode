package day04

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io"

	"github.com/1e9y/adventofcode/challenge"
)

func check(h hash.Hash, condition string) bool {
	s := hex.EncodeToString(h.Sum(nil))
	if s[:len(condition)] == condition {
		return true
	}
	return false
}

func A(input *challenge.Challenge) int {
	secret := input.LineSlice()[0]
	answer := 0
	condition := "00000"

	var h hash.Hash
	for {
		h = md5.New()
		_, err := io.WriteString(h, fmt.Sprintf("%s%d", secret, answer))
		if err != nil {
			panic(err)
		}

		if check(h, condition) {
			return answer
		}

		answer++
	}

	return answer
}

func B(input *challenge.Challenge) int {
	secret := input.LineSlice()[0]
	condition := "000000"

	answer := make(chan int)
	for i := 0; i < 4; i++ {
		a := i
		go func() {
			var h hash.Hash
			for {
				h = md5.New()
				_, err := io.WriteString(h, fmt.Sprintf("%s%d", secret, a))
				if err != nil {
					panic(err)
				}

				if check(h, condition) {
					answer <- a
					return
				}

				a += 4
			}
		}()
	}

	return <-answer
}
