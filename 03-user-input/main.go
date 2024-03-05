package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	singleQuote()
	readUserInput()
}

func singleQuote() {
	// char within single quote represents a rune literal
	// Its type is int32, and you can think them as char in Java
	newLineChar := '\n'

	// double quote returns a string
	newLineString := "\n"
	fmt.Printf("Variable is of type: %T \n", newLineChar)
	fmt.Printf("Variable is of type: %T \n", newLineString)
}

func readUserInput() {
	fmt.Printf("Please input a value: ")

	reader := bufio.NewReader(os.Stdin)

	// Below format is call comma ok/err format
	// Equivalent to try/catch in Java
	// If you don't use the value of err, then use "_"
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your input: ", input)
	fmt.Printf("Variable is of type: %T \n", input)
}
