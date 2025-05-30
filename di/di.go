package di

import (
	"bytes"
	"fmt"
)

// passing a buffer to Fprintf allows capturing the
// result of outputting to stdout, making this func
// testable
func Greet(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}