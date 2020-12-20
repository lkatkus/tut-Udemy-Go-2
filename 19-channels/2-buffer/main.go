package main

import "fmt"

func main() {
	// (type of channel, buffer size)
	c := make(chan int, 2)

	c <- 42
	c <- 54

	fmt.Println(<-c)
	fmt.Println(<-c)
}

// Unsuccesfull buffer
// func main() {
// 	c := make(chan int, 1)

// 	c <- 42
// 	c <- 43

// 	fmt.Println(<-c)
// }
