package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 0; i < 128; i++ {
		fmt.Printf("%7b %#4o  %4d  %#4X  %6U  %3s  %s \n", i, i, i, i, i, strconv.Itoa(i), string(i))
	}
}
