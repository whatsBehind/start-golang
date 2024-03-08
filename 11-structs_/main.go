package main

import "fmt"

func main() {
	printUser()
}

func createAUser() User {
	return User{"whatsbehind", "whatsbehind@gmail.com", 18, true}
}

func printUser() {
	var user = createAUser()
	fmt.Println(user)
	// use "+" to print user instance with field name and value
	fmt.Printf("Detail about a user %+v", user)
}

type User struct {
	Name  string
	Email string
	Age   int
	Login bool
}
