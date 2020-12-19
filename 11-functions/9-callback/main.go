package main

import "fmt"

func main() {
	foo(bar)
}

func foo(cb func()) {
	fmt.Println("Inside foo")

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	cb()
}

func bar() {
	fmt.Println("Inside bar")
}
