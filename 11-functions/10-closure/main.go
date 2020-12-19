package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		foo(i)
	}

	{
		y := 42
		fmt.Println(y)
	}

	// y no longer available
	// fmt.Println(y)

	a := incrementor()
	b := incrementor()

	fmt.Println("--- A ---")
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println("--- B ---")
	fmt.Println(b())
	fmt.Println(b())

	addToFour := add(4)
	addToTen := add(10)

	fmt.Println("--- add ---")
	addToFour(2)
	addToFour(99)
	addToTen(1)
	add(99)(1)
}

func foo(i int) {
	fmt.Println(i)
}

func incrementor() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}

func add(x int) func(y int) int {
	return func(y int) int {
		sum := x + y
		fmt.Println(sum)

		return sum
	}
}
