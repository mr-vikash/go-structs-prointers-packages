package main

import "fmt"

type user struct {
	name     string
	email    string
	password string
	age      int
}

func main() {
	var u1 user

	u1.age = 22
	u1.email = "mrvikash@example.com"
	u1.password = "password@123"
	u1.name = "Vikash Gupta"

	fmt.Println("The First User is: ", u1)

	var u2 user
	u2.age = 25
	u2.email = "mrankur@example.com"
	u2.password = "password@2234"
	u2.name = "Ankur Gupta"

	fmt.Println("The Second User is: ", u2)
}
