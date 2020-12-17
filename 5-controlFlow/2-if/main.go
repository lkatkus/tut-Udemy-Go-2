package main

import "fmt"

func main() {
	if true {
		fmt.Println("Well, it's true")
	}

	if false {
		fmt.Println("What did you expect?")
	}

	if !false {
		fmt.Println("Yes, it does")
	}

	if 2 == 2 {
		fmt.Println("I hope so")
	}

	if !(2 != 2) {
		fmt.Println("Do you really need this?")
	}

	// initialization statement and condition
	if x := 42; x == 42 {
		fmt.Println("what..?", x)
	}

	if false {
		fmt.Println("Nope")
	} else if false {
		fmt.Println("Still nope")
	} else {
		fmt.Println("Finally!")
	}

	evenOdd()
}

func evenOdd() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Printf("%v is even\n", i)
		} else {
			fmt.Printf("%v is odd\n", i)
		}
	}
}
