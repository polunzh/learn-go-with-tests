package iteration

import (
	"strings"
)

func Repeat(character string, times int) string {
	var sb strings.Builder

	for i := 0; i < times; i++ {
		sb.WriteString(character)
	}

	return sb.String()
}
