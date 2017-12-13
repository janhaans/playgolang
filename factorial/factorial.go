package main

import (
	"fmt"
)

func main() {
	for f := range factorial(gen()) {
		fmt.Println(f)
	}
}

func gen() chan int {
	c := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		for in := range c {
			result := 1
			for i := 1; i <= in; i++ {
				result *= i
			}
			out <- result
		}
		close(out)
	}()
	return out
}
