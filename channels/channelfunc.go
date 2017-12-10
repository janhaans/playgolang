package main

import (
	"fmt"
)

func ProduceText() chan string {
	c := make(chan string)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- "Hello"
			//time.Sleep(100 * time.Millisecond)
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- "Ciao"
			//time.Sleep(100 * time.Millisecond)
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	return c

}

func main() {
	for t := range ProduceText() {
		fmt.Println(t)
	}
}
