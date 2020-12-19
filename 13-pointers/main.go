package main

import "fmt"

func main() {
	fmt.Println("--- 1 ---")
	initial()
	fmt.Println("--- 2 ---")
	x := 11
	foo(x)
	fmt.Println("x", x)
	bar(&x)
	fmt.Println("x", x, &x)
	fmt.Println("--- 3 ---")
}

func initial() {
	a := 42

	fmt.Println(a)         // a value
	fmt.Println(&a)        // pointer to a
	fmt.Printf("%T\n", a)  // int - an int
	fmt.Printf("%T\n", &a) // *int - a pointer to an int

	b := &a
	fmt.Println(b)
	fmt.Println(*b)  // returns value stored at the location pointed at
	fmt.Println(*&a) // * gives an adress and & gets the value stored at the address

	*b = 43
	fmt.Println("b", b)
	fmt.Println("a", a)
}

func foo(y int) {
	fmt.Println(y)
	y = 43
	fmt.Println(y)
}

func bar(y *int) {
	fmt.Println(y, *y)
	*y = 99
}
