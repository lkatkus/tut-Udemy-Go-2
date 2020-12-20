package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("NumCPU:", runtime.NumCPU())
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	counter := 0
	const gs = 10
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("NumGoroutine after:", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
