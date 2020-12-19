package main

import "fmt"

func main() {
	f := foo()
	b := bar()

	f()
	fmt.Println(b)
}

func foo() func() {
	fmt.Println("inside foo")

	return func() {
		fmt.Println("from function returned by foo")
	}
}

func bar() string {
	s := "inside bar"

	return s
}
