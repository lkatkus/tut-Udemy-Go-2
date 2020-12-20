package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("Runs before func main()")
}

var wg sync.WaitGroup

func main() {
	fmt.Println("GOOS\t\t", runtime.GOOS)
	fmt.Println("GOARCH\t\t", runtime.GOARCH)
	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("NumGoroutine\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("NumGoroutine\t", runtime.NumGoroutine())

	wg.Wait()

	ch := make(chan int)

	go func() {
		ch <- doSomething(3)
	}()

	fmt.Println("Data passed through a channel:", <-ch)
}

func foo() {
	fmt.Println("--- FOO ---")

	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}

	wg.Done()
}

func bar() {
	fmt.Println("--- BAR ---")

	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}

func doSomething(x int) int {
	return x * 5
}
