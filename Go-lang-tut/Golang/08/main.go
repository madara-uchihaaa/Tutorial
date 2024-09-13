package main

import "fmt"

func main() {
	// Structs in Golang

	type User struct {
		// U can also use `user` instead of `User` but it is not recommended
		ID        int
		FirstName string
		LastName  string
		Email     string
		IsActive  bool

		// I not i in `IsActive` because it is public
	}

	var user1 User
	// fmt.Println(user1)

	user1.ID = 1
	user1.FirstName = "John"
	user1.LastName = "Doe"
	user1.Email = "risha@gmail.com"
	user1.IsActive = true

	fmt.Println(user1)

	user2 := User{
		ID:        2,
		FirstName: "Jane",
		LastName:  "Doe",
		Email:     "x@x.com",
		IsActive:  false,
	}

	fmt.Println(user2)
	fmt.Printf("User2: %v\n User: %+v\n", user2,user1)

}
