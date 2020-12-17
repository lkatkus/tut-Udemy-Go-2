package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	var x = 4
	fmt.Printf("Decimal: %d\tBinary: %b\n", x, x)

	var y = x >> 1
	fmt.Printf("Decimal: %d\tBinary: %b\n", y, y)

	var z = x << 1
	fmt.Printf("Decimal: %d\tBinary: %b\n", z, z)

	// kb := 1024
	// mb := 1024 * kb
	// gb := 1024 * mb

	fmt.Printf("Decimal: %d\t\tBinary: %b\n", kb, kb)
	fmt.Printf("Decimal: %d\tBinary: %b\n", mb, mb)
	fmt.Printf("Decimal: %d\tBinary: %b\n", gb, gb)
}
