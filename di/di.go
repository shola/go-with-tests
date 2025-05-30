package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Greet(os.Stdout, "Elodie")
}

// passing an io.Writer interface param (includes buffer and stdout) to Fprintf allows 
// capturing the result of outputting to stdout, making this func testable
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}