package main

import "fmt"

func main() {
	ids := []int{46, 68, 97, 45, 36, 52, 19}

	// Loop though ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Exclude index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Mike": "mike@gmail.com", "Jenny": "jenny@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

}
