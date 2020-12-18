package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	var y [6]int

	// Array length is part of array type
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	x[3] = 42
	fmt.Println(x)
}
