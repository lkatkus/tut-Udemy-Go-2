package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("func expression")
	}

	e := func(i int) {
		fmt.Println("func expression with arguments", i)
	}

	f()
	e(1984)
}
