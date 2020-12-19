package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous function call")
	}()

	func(s string) {
		fmt.Println("Anonymous function call with params:", s)
	}("Anonymous argument")

	foo(func() {
		fmt.Println("bar")
	})
}

func foo(cb func()) {
	fmt.Println("foo")
	cb()
}
