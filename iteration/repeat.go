package iteration

import (
	"strings"
)

// Building instead of concatenating strings minimizes copying data
// for new immutable strings
func Repeat(char string, length int) string {
	var repeated strings.Builder

	for i := 0; i < length; i++ {
		repeated.WriteString(char)
	}

	return repeated.String()
}
