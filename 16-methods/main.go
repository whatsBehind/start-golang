package main

import "fmt"

func main() {
	var user = createAUser()
	user.IsLogin()
	fmt.Printf("Detail about the user %+v", user)
	user.SetEmail("test@gmail.com")
	fmt.Printf("Detail about the user %+v", user)
}

func createAUser() User {
	return User{"whatsbehind", "whatsbehind@gmail.com", 18, true}
}

type User struct {
	Name  string
	Email string
	Age   int
	Login bool
}

func (user User) IsLogin() {
	fmt.Println("Is user login?", user.Login)
}

// The user passed in is not the original instance but a copy
// In other words, the set method won't modify the original instance
func (user User) SetEmail(email string) {
	user.Email = email
}
