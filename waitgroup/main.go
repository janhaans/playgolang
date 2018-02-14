package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayHello(name string, w *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("Hello %s \n", name)
	}
	w.Done()
}

func main() {
	wg.Add(2)
	fmt.Println("The test begins")
	go SayHello("Jan", &wg)
	go SayHello("AM", &wg)
	wg.Wait()
	fmt.Println("The test ends")
}
