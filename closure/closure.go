package main

import (
	"fmt"

	"github.com/janhaans/playgolang/counter"
)

func main() {
	start1 := 10
	start2 := 20
	fmt.Printf("start:\t\tcounter1 = %d\tcounter2 = %d\n", start1, start2)
	increment1 := counter.IncrementCounter(start1)
	increment2 := counter.IncrementCounter(start2)

	for i := 1; i <= 10; i++ {
		fmt.Printf("Iteration %d:\tcounter1 = %d\tcounter2 = %d\n", i, increment1(), increment2())
	}
}
