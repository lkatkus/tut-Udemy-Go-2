package main

import "fmt"

func main() {
	foo("first")
	fmt.Println("------------")
	foo("first", "second")
	fmt.Println("------------")
	foo("first", "second", "third")
	fmt.Println("------------")
	sum := bar(1, 2, 3)
	fmt.Println(sum)
	sum = bar(1, 2, 3, 4, 5)
	fmt.Println(sum)
	sum = bar(99, 1)
	fmt.Println(sum)

	xi := []int{1, 2, 3, 4, 5}
	sum = bar(xi...)
	fmt.Println(sum)
	// variadic 0 or more
	sum = bar()
	fmt.Println(sum)
}

func foo(s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func bar(s ...int) int {
	sum := 0

	for _, v := range s {
		sum = sum + v
	}

	return sum
}
