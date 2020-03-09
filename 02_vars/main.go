package main

import "fmt"

func main() {

	// MAIN TYPES
	// string
	// bool
	// int
	// uint (no negative numbers)
	// byte - alias fir uint32
	// rune - alias for int32
	// float34 float64
	// complex64 complex128

	// var name = "Mike"
	var age int = 45
	const isCool = true

	// Shorthand declaration
	// name := "Mike"
	size := 1.5
	name, email := "Mike", "mike@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("Type: %T", size)

}
