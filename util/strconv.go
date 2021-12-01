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
