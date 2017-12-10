package main

import (
	"fmt"
)

func main() {
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
		for t := range c {
			fmt.Println(t, " Jan")
			//time.Sleep(100 * time.Millisecond)
		}
		done <- true
	}()

	go func() {
		for t := range c {
			fmt.Println(t, " AM")
			//time.Sleep(time.Millisecond)
		}
		done <- true
	}()

	<-done
	<-done
	close(c)
	<-done
	<-done

}
