package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("--- foo ---")
	foo()

	fmt.Println("--- bar ---")
	bar()

	fmt.Println("--- nextFoo ---")
	nextFoo()

	fmt.Println("--- BEFORE EXIT ---")
}

func foo() {
	// bi-directional channel
	c := make(chan int)
	// send-only channel
	cs := make(chan<- int)
	// receive-only channel
	cr := make(<-chan int)

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cs)
	fmt.Printf("%T\n", cr)
}

func bar() {
	c := make(chan int)

	go send(c)
	receive(c)
}

func send(c chan<- int) {
	time.Sleep(time.Second * 3)
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}

func nextFoo() {
	c := make(chan int)

	go func(c chan<- int) {
		for i := 0; i < 10; i++ {
			c <- i
		}

		close(c)
	}(c)

	for v := range c {
		fmt.Println(v)
	}
}
