package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 123
	y := 22.111
	z := float64(4 / 3)

	fmt.Printf("%v\t%T\n", x, x)
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Printf("%v\t%T\n", z, z)

	fmt.Println(runtime.GOARCH, runtime.GOOS)
}
