package main

import "fmt"

func main() {
	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("a type: %T\nb type: %T\n", a, b)

	// Use * to read val from address
	fmt.Println(*b)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
