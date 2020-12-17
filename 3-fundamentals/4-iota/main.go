package main

import "fmt"

const (
	a = iota + 1
	b
	c
)

const (
	d = iota
	e
	f
)

func main() {
	fmt.Printf("%T\t%v\n", a, a)
	fmt.Printf("%T\t%v\n", b, b)
	fmt.Printf("%T\t%v\n", c, c)

	fmt.Printf("%T\t%v\n", d, d)
	fmt.Printf("%T\t%v\n", e, e)
	fmt.Printf("%T\t%v\n", f, f)
}
