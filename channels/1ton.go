package main

import "fmt"

func main() {
	c := make(chan string)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- "Hello"
		}
		close(c)
	}()

	go func() {
		for t := range c {
			fmt.Println(t, " Jan")
		}
		done <- true
	}()

	go func() {
		for t := range c {
			fmt.Println(t, " AM")
		}
		done <- true
	}()

	<-done
	<-done

}
