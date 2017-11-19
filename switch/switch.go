package main

import "fmt"

//Type switch
func do(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float"
	case bool:
		return "bool"
	default:
		return "undetermined type"
	}
}

func main() {
	var b bool
	var i int
	var f float64
	var s string
	fmt.Printf("Type of variable b = %s\n", do(b))
	fmt.Printf("Type of variable i = %s\n", do(i))
	fmt.Printf("Type of variable f = %s\n", do(f))
	fmt.Printf("Type of variable s = %s\n", do(s))
}
