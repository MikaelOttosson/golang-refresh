package main

import "fmt"

func greeting(name string) string {
	return "Greetings, " + name
}

func getSum(num1, num2 int) int {
	return num1 * num2

}

func main() {
	fmt.Println(greeting("Mike"))
	fmt.Println("Sum:", getSum(54, 34))

}
