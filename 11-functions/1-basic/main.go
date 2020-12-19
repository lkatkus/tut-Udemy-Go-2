package main

import "fmt"

func main() {
	foo()
	bar("VERY_STRING_MUCH_ARGUMENT")
	fmt.Println(woo("zoo"))
	x, y := hoo("outside hoo")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Printed string", s)
}

func woo(s string) string {
	return fmt.Sprint("Returned string ", s)
}

func hoo(s string) (string, string) {
	return fmt.Sprint("First return value"), fmt.Sprint("Your passed value as second return value ", s)
}
