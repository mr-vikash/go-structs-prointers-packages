package main

import (
	"fmt"

	"example.com/struct/user"
)

func main() {
	var u1 user.User

	u1.Age = 22
	u1.Email = "mrvikash@example.com"
	u1.Password = "password@123"
	u1.Name = "Vikash Gupta"

	fmt.Println("The First User is: ", u1)

	var u2 user.User
	u2.Age = 25
	u2.Email = "mrankur@example.com"
	u2.Password = "password@2234"
	u2.Name = "Ankur Gupta"

	fmt.Println("The Second User is: ", u2)
}
