package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
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
}

func task1() {
	var wg sync.WaitGroup

	wg.Add(2)

	fmt.Println("Before goroutines")

	go func() {
		fmt.Println("foo")
		wg.Done()
	}()
	go func() {
		fmt.Println("bar")
		wg.Done()
	}()

	fmt.Println("After goroutines")

	wg.Wait()

	fmt.Println("After wait")
}

type person struct {
	first, last string
}

func (p *person) speak() {
	fmt.Println((*p).first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func task2() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}

	saySomething(&p1)
	// does not work
	// saySomething(p2)
	p2.speak()
}

func task3() {
	counter := 0
	times := 420

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			mu.Lock()
			runtime.Gosched()
			v := counter
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Counter is", counter)
}

func task4() {
	var counter int64 = 0
	times := 9999

	var wg sync.WaitGroup
	wg.Add(times)

	for i := 0; i < times; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Counter is", counter)
}

func task5() {
	fmt.Println("GOOS", runtime.GOOS)
	fmt.Println("GOARCH", runtime.GOARCH)
	fmt.Println("NumCPU", runtime.NumCPU())
}
