package main

import (
	"fmt"
	"sync"
)

var counter int64
var wg sync.WaitGroup

//This program has a no race condition due to mutex
//Check that there is no race condition with go run -race main.go
func main() {
	mutex := sync.Mutex{}
	counter = 0
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			mutex.Lock()
			counter++
			mutex.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Printf("Counter should be 20 and is %d \n", counter)

}
