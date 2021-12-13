package day08

import (
	"fmt"
	"strings"

	"github.com/1e9y/adventofcode/challenge"
)

/*
  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
*/

func parseInput(input <-chan string) (result [][][]string) {
	for line := range input {
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			panic(fmt.Sprintf("bad input: %s", line))
		}

		r := make([][]string, 2)
		for i, part := range parts {
			r[i] = strings.Split(part, " ")
		}

		result = append(result, r)
	}
	return
}

func appearance(input <-chan string) (result int) {
	parsedInput := parseInput(input)
	for _, signal := range parsedInput {
		for _, digit := range signal[1] {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				result++
			}
		}
	}
	return
}

// func break(signals []string) {

// }

func equal(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[rune]bool)
	for _, r := range a {
		m[r] = true
	}
	for _, r := range b {
		if !m[r] {
			return false
		}
	}
	return true
}

func includes(a, b string) bool {
	m := make(map[rune]bool)
	for _, r := range a {
		m[r] = true
	}
	for _, r := range b {
		if !m[r] {
			return false
		}
	}
	return true
}

func add(a, b string) string {
	m := make(map[rune]bool)
	for _, r := range a {
		m[r] = true
	}

	for _, r := range b {
		m[r] = true
	}

	result := ""
	for c, ok := range m {
		if ok {
			result += string(c)
		}
	}

	return result
}

func subtract(a, b string) string {
	m := make(map[rune]bool)
	for _, r := range a {
		m[r] = true
	}

	for _, r := range b {
		m[r] = false
	}

	result := ""
	for c, ok := range m {
		if ok {
			result += string(c)
		}
	}

	return result
}

func decode(input <-chan string) (result int) {
	parsedInput := parseInput(input)

	// testLine := parsedInput[0]
	for i, testLine := range parsedInput {
		if i != 0 {
			// continue
		}
		// fmt.Println(testLine)
		signals := testLine[0]
		inputs := testLine[1]

		codesByLength := make(map[int][]string)
		codeToN := make(map[string]int)
		nToCode := make(map[int]string)

		for _, signal := range signals {
			l := len(signal)
			if _, ok := codesByLength[l]; !ok {
				codesByLength[l] = make([]string, 0)
			}
			codesByLength[l] = append(codesByLength[l], signal)
		}

		// 1
		code1 := codesByLength[2][0]
		nToCode[1] = code1
		codeToN[code1] = 1
		// fmt.Printf("code %v is %v\n", 1, code1)

		// 4
		code4 := codesByLength[4][0]
		nToCode[4] = code4
		codeToN[code4] = 4
		// fmt.Printf("code %v is %v\n", 4, code4)

		// 7
		code7 := codesByLength[3][0]
		nToCode[7] = code7
		codeToN[code7] = 7
		// fmt.Printf("code %v is %v\n", 7, code7)

		// 8
		code8 := codesByLength[7][0]
		nToCode[8] = code8
		codeToN[code8] = 8
		// fmt.Printf("code %v is %v\n", 8, code8)

		var code2 string
		var tryCode string
		// 2
		tryCode = subtract(code8, add(code7, code4))
		// println(tryCode)
		for _, s := range codesByLength[5] {
			if includes(s, tryCode) {
				codeToN[s] = 2
				nToCode[2] = s
				code2 = s
			}
		}

		// fmt.Printf("code %v is %v\n", 2, code2)
		// 9
		var code9 string
		for _, s := range codesByLength[6] {
			if !includes(s, tryCode) {
				// println(s, tryCode)
				codeToN[s] = 9
				nToCode[9] = s
				code9 = s
			}
		}
		// fmt.Printf("code %v is %v\n", 9, code9)
		// // 6
		// for _, s := range codesByLength[6] {
		// 	if includes(s, tryCode) {
		// 		// println(s, tryCode)
		// 		codeToN[s] = 6
		// 		nToCode[6] = s
		// 		code6 = s
		// 	}
		// }
		// fmt.Printf("code %v is %v\n", 6, code6)

		// 5
		var code5 string
		tryCode = subtract(code9, add(code7, code2))
		for _, s := range codesByLength[5] {
			if includes(s, tryCode) {
				codeToN[s] = 5
				nToCode[5] = s
				code5 = s
			}
			// fmt.Printf("has %s %s? %v\n", s, somecode, includes(s, somecode))
		}

		// 3
		// var code3 string
		for _, s := range codesByLength[5] {
			if s != code5 && s != code2 {
				codeToN[s] = 3
				nToCode[3] = s
				// code3 = s
			}
			// fmt.Printf("has %s %s? %v\n", s, somecode, includes(s, somecode))
		}

		// 0
		// 6
		// var code0 string
		// var code6 string
		for _, s := range codesByLength[6] {
			if codeToN[s] == 9 {
				continue
			}
			tryCode = subtract(s, code5)
			if len(tryCode) == 1 {
				codeToN[s] = 6
				nToCode[6] = s
				// code6 = s
			}
			// fmt.Printf("has %s %s? %v\n", s, somecode, includes(s, somecode))
			if len(tryCode) == 2 {
				codeToN[s] = 0
				nToCode[0] = s
				// code0 = s
			}
		}

		var output int
		for _, s := range inputs {
			var d int
			for code, n := range codeToN {
				if equal(s, code) {
					d = n
					// break
				}
			}
			output += d
			output *= 10
		}
		output /= 10
		result += output
		// fmt.Println(output)

		// fmt.Println(codeToN)
		// fmt.Println(nToCode)
	}

	return
}

func A(input *challenge.Challenge) int {
	return appearance(input.Lines())
}

func B(input *challenge.Challenge) int {
	return decode(input.Lines())
}
