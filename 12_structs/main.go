package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName string
	lastName  string
	age       int
}

// Greeting method (value reciever) and cast a int to string with Iota
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and my age is " + strconv.Itoa(p.age)
}

// Pointer reciever method
func (p *Person) hasBirthday() {
	p.age++
}

// Pointer reciever method
func (p *Person) changedLastName(newLastName string) {
	p.lastName = newLastName + "(new lastname)"
}

func main() {
	// //init person
	person1 := Person{firstName: "Michael", lastName: "Ottosson", age: 45}
	person2 := Person{"Jenny", "Svensson", 45}

	// fmt.Println(person1.firstName, person2.age)

	person1.hasBirthday()
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
	person2.changedLastName("Ottosson")
	fmt.Println(person2.greet())
}
