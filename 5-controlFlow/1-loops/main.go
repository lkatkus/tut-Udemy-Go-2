package main

import "fmt"

func main() {
	// for init; condition; post {
	// 	action
	// }

	for i := 0; i <= 5; i++ {
		fmt.Println("Outter", i)

		for j := 0; j <= 5; j++ {
			fmt.Println("\tInner", j)
		}
	}

	x := 1

	for x < 5 {
		fmt.Println("x", x)
		x++
	}

	for {
		if x > 10 {
			break
		}

		fmt.Println("while loop x", x)

		x++
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Println("Even:", i)
	}

	printASCII()
}

func printASCII() {
	for i := 33; i < 122; i++ {
		fmt.Printf("%c - %#x\n", i, i)
	}
}
