package main

import "fmt"

// main function is the entry point of your program
// You don't need to invoke main function.
// Golang will find the main function and invoke it
func main() {
	var result = add(1, 2)
	fmt.Printf("The result is %v\n", result)

	result = addAll(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("The result is %v\n", result)

	number, value := returnMoreThanOne()
	fmt.Printf("The number is %v, and the string is %v\n", number, value)
}

func add(n1 int, n2 int) int {
	return n1 + n2
}

// values here is a slice of int
func addAll(values ...int) int {
	result := 0
	for _, value := range values {
		result += value
	}
	return result
}

func returnMoreThanOne() (int, string) {
	return 0, "Hello World"
}
