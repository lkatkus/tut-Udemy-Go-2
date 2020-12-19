package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello, World")
	io.WriteString(os.Stdout, "Hello, WriteString")
}
