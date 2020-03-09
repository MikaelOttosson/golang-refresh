package main

import "fmt"

func main() {
	// maps - key value pairs
	// emails := make(map[string]string)

	// emails["Mike"] = "mike@gmail.com"
	// emails["Jenny"] = "jenny@gmail.com"

	emails := map[string]string{"Mike": "mike@gmail.com", "Jenny": "jenny@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Mike"])

	// Delete ftom maps
	delete(emails, "Jenny")
	fmt.Println(emails)
	fmt.Println(len(emails))
}
