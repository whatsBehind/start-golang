package main

import "fmt"

// Below usage is not allowed
// globalIntNum := 100

// to declare a global variable, you have to use key word var
// The scope of variable starting with lower case is this module only
// which is equivalent to private in Java
var globalBoolVal bool = false

// public global variable starting with upper case letter
var GlobalBoolVal bool = true

func main() {
	printType()
	varDeclaration()
}

func varDeclaration() {
	var userName string = "whatsbehind"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n", userName)

	// := definition, no key word var is required
	// := can only be used inside a function
	intNum := 100
	fmt.Println(intNum)
	fmt.Printf("Variable is of type: %T \n", intNum)

	fmt.Println(globalBoolVal)
	fmt.Printf("Variable is of type: %T \n", globalBoolVal)

	fmt.Println(GlobalBoolVal)
	fmt.Printf("Variable is of type: %T \n", GlobalBoolVal)

	// implicit type
	var implicitIntNum = 100
	fmt.Println(implicitIntNum)
	fmt.Printf("Variable is of type: %T \n", implicitIntNum)
}

func printType() {
	var userName string = "whatsbehind"

	// `%T` format verb is only specific to printf
	fmt.Printf("Variable is of type: %T \n", userName)
	fmt.Println("Variable is of type: %T \n", userName)
}
