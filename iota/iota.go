package main

import "fmt"

const (
	zero = iota
	one
	two
	three
)

func main() {
	fmt.Printf("Zero:\t%4b\t%1d\n", zero, zero)
	fmt.Printf("One:\t%4b\t%1d\n", one, one)
	fmt.Printf("Two:\t%4b\t%1d\n", two, two)
	fmt.Printf("Three:\t%4b\t%1d\n", three, three)
}
