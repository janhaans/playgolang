package main

import "fmt"

func Filter(ints []int, callback func(i int) bool) []int {
	var xs []int
	for _, v := range ints {
		if callback(v) {
			xs = append(xs, v)
		}
	}
	return xs
}

func Even(i int) bool {
	if i%2 == 0 {
		return true
	} else {
		return false
	}
}

func Uneven(i int) bool {
	if i%2 == 0 {
		return false
	} else {
		return true
	}
}

func main() {
	reeks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Even values are: %v\n", Filter(reeks, Even))
	fmt.Printf("Uneven values are: %v\n", Filter(reeks, Uneven))
}
