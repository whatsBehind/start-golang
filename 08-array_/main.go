package main

import "fmt"

func main() {
	createArray()
}

func createArray() {
	var myArray [4]string
	myArray[0] = "Hello"
	myArray[1] = "WhatsBehind"
	myArray[2] = "!"

	fmt.Println("My array: ", myArray)

	var initiatedArray = [4]string{"Hello", "World", "!"}
	fmt.Println("Initiated array: ", initiatedArray)
	fmt.Println("Length of initiated array: ", len(initiatedArray))
}
