package main

import (
	"fmt"
)

func main() {
	fmt.Println("------ Task 1 ------")
	task1()
	fmt.Println("------ Task 2 ------")
	task2()
	fmt.Println("------ Task 3 ------")
	task3()
	fmt.Println("------ Task 4 ------")
	task4()
	fmt.Println("------ Task 5 ------")
	task5()
	fmt.Println("------ Task 6 ------")
	task6()
	fmt.Println("------ Task 7 ------")
	task7()
}

func task1() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	cb := make(chan int, 1)
	cb <- 22

	fmt.Println(<-cb)
}

func task2() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println(cs)
	fmt.Println(<-cs)
	fmt.Printf("cs\t%T\n", cs)
}

func task3() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func task4() {
	q := make(chan int)
	c := genTask4(q)

	receiveTask4(c, q)

	fmt.Println("about to exit")
}

func genTask4(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		q <- 0
		close(c)
	}()

	return c
}

func receiveTask4(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("value", v)
		case v := <-q:
			fmt.Println("quit", v)
			return
		}
	}
}

func task5() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func task6() {
	c := make(chan int)

	go getNumbers(c)
	readNumbers(c)

	fmt.Println("Exit task6")
}

func getNumbers(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}

	close(c)
}

func readNumbers(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}

func task7() {
	c := make(chan string)

	for i := 0; i < 10; i++ {
		go func(c2 chan<- string, numb int) {
			for j := 0; j < 10; j++ {
				value := fmt.Sprintf("%v-%v", numb, j)
				c2 <- value
			}

		}(c, i)
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k)
	}
}
