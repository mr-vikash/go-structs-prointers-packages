package user

import "errors"

type User struct {
	Name     string
	Email    string
	Password string
	Age      int
}

func NewUser(name string, email string, password string, age int) (*User, error) {
	if name == "" || email == "" || password == "" {
		return nil, errors.New("Please Enter the values of required fields")
	}
	return &User{
		Name:     name,
		Email:    email,
		Password: password,
		Age:      age,
	}, nil
}

func (u User) IsAdult(age int) bool {
	if age >= 18 {
		return true
	}
	return false
}
