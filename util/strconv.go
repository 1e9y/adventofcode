package util

import (
	"fmt"
	"strconv"
)

func MustAtoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(fmt.Errorf("util.MustAtoi: %w", err))
	}
	return i
}

func MustParseInt(a string, base int) int {
	i, err := strconv.ParseInt(a, base, 0)
	if err != nil {
		panic(fmt.Errorf("util.MustParseInt: %w", err))
	}
	return int(i)
}
