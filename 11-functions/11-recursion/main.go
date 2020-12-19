package main

import "fmt"

func main() {
	n := factorial(4)
	fmt.Println(n)

	m := factorial(13)
	fmt.Println(m)

	l := loopFactorial(4)
	fmt.Println(l)
}

func factorial(n int) int {
	fmt.Println("Current", n)

	if n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	total := 1

	for i := n; i > 0; i-- {
		total *= i
	}

	return total
}
