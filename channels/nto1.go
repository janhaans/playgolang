package main

import "fmt"

func main() {
	c := make(chan string)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- "Hello"
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			c <- "Ciao"
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for t := range c {
		fmt.Println(t)
	}

}
