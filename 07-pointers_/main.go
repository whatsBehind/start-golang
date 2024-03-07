package main

import "fmt"

func main() {
	printPointer()
}

func printPointer() {
	myNumber := 23
	// Create a pointer which is a reference to myNumber, same memory
	myPointer := &myNumber

	fmt.Println("The value of the pointer: ", myPointer)
	// Use "*" to get the value of the variable pointer pointing to
	fmt.Println("Value variable the pointer pointing to: ", *myPointer)

	*myPointer += 2
	fmt.Println("Value of myNumber", myNumber)
}
