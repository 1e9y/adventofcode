package util

import (
	"fmt"
	"strings"
)

func Key(n ...int) string {
	list := []string{}
	for _, a := range n {
		list = append(list, fmt.Sprintf("%d", a))
	}
	return strings.Join(list, ":")
}
