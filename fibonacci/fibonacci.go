package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		//	fmt.Print("1 ")
		return 1
	}
	//fmt.Printf("%d ", Fibonacci(n-1)+Fibonacci(n-2))
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	var input int
	fmt.Print("Give a number: ")
	fmt.Scanf("%d", &input)
	fmt.Printf("Fibonacci sequenc of %d = %d\n", input, Fibonacci(input))
	fmt.Println("")
}
