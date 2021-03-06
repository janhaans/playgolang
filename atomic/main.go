package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

//This program has a race condition
//go run -race shows that this program has a race condition
func main() {
	counter = 0
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			atomic.AddInt64(&counter, 1)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			atomic.AddInt64(&counter, 1)
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("Counter should be 20 and is %d \n", counter)

}
