package main

import "fmt"

func main() {
	printDefer()
}

// defer statements will be cut and executed right before "}"
// When there are multiple defer statements, their execution order is reversed
func printDefer() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	// All defer statements will be executed after invoking and finishing all functions
	nestedDefer()
	fmt.Println("Hello World")
}

func nestedDefer() {
	defer fmt.Println("nested defer statement")
}
