package main

import "fmt"

func main() {
	// Arrays has fixed lengths and have fixed types
	// var fruitArr [2]string

	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Better Faster formatting
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape"}

	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:])

}
